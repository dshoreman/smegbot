package commands

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/config"
	"github.com/dshoreman/smegbot/util"
)

func nuke(s *dg.Session, m *dg.MessageCreate) {
	config.LoadGuild(m.GuildID)
	sinbin := quarantineRole(s, m.ChannelID, m.GuildID)
	if sinbin != "" {
		u := m.Mentions[0].ID
		err := replaceRoles(s, m.GuildID, m.ChannelID, u, sinbin)
		sendSuccessOrFail(s, m.ChannelID, err, "add", u)
	}
}

func replaceRoles(s *dg.Session, g string, c string, u string, sinbin string) error {
	roles := memberRoles(s, g, u)
	if err := s.GuildMemberRoleAdd(g, u, sinbin); err != nil {
		fmt.Println("Could not add quarantine role", err)
		return err
	}
	s.ChannelMessageSend(c, "Please wait while the member's roles are removed...")
	err := util.WriteJSON(util.GuildPath("m.roles", g, u), roles)
	if err == nil {
		removeRoles(s, g, u, roles)
	}
	return err
}

func restore(s *dg.Session, m *dg.MessageCreate) {
	sinbin := quarantineRole(s, m.ChannelID, m.GuildID)
	if sinbin != "" {
		target := m.Mentions[0].ID
		restoreRoles(s, m.GuildID, m.ChannelID, target)

		err := s.GuildMemberRoleRemove(m.GuildID, target, sinbin)
		sendSuccessOrFail(s, m.ChannelID, err, "remove", target)
	}
}

func quarantineRole(s *dg.Session, channelID string, guildID string) string {
	roles, _ := s.GuildRoles(guildID)
	if config.Guild.QuarantineRole != "" {
		return config.Guild.QuarantineRole
	}
	for _, r := range roles {
		if r.Name == "Quarantine" {
			return r.ID
		}
	}
	s.ChannelMessageSend(channelID, "I couldn't find the **@Quarantine** role!")
	return ""
}

func memberRoles(s *dg.Session, guildID string, target string) []string {
	member, err := s.GuildMember(guildID, target)
	roles := make([]string, 0)
	if err == nil {
		roles = append(roles, member.Roles...)
	}
	return roles
}

func removeRoles(s *dg.Session, g string, u string, roles []string) {
	for _, role := range roles {
		fmt.Printf("Removing role %s...\n", role)
		s.GuildMemberRoleRemove(g, u, role)
	}
}

func restoreRoles(s *dg.Session, g string, c string, u string) {
	roles := make([]string, 0)
	if util.ReadJSON(util.GuildPath("m.roles", g, u), &roles) != nil {
		return
	}
	s.ChannelMessageSend(c, "Please wait, role restoration can take a while...")
	for _, role := range roles {
		s.GuildMemberRoleAdd(g, u, role)
	}
}

func sendSuccessOrFail(s *dg.Session, channelID string, err error, mode string, target string) {
	op, result := "add", "now in Quarantine."
	if mode == "remove" {
		op, result = "remove", "back out of Quarantine!"
	}
	if err != nil {
		s.ChannelMessageSend(channelID, "Oops! Couldn't "+op+" the configured quarantine role. Check permissions!")
		fmt.Println("\nError:\n", err)
		return
	}
	s.ChannelMessageSend(channelID, fmt.Sprintf("<@%s> is %s", target, result))
}
