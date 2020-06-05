package commands

import (
	"fmt"

	discord "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/util"
)

func nuke(s *discord.Session, m *discord.MessageCreate) {
	sinbin := quarantineRole(s, m.ChannelID, m.GuildID)
	if sinbin != "" {
		g, u, saved := m.GuildID, m.Mentions[0].ID, true
		roles := memberRoles(s, g, u)

		err := util.WriteJSON(util.GuildPath("m.roles", g, u), roles)
		if err != nil {
			fmt.Println("\nError: Could not save roles to file\n", err)
			saved = false
		}
		err = s.GuildMemberRoleAdd(g, u, sinbin)
		if saved && err == nil {
			removeRoles(s, g, u, roles)
		}
		sendSuccessOrFail(s, m.ChannelID, err, "add", u)
	}
}

func restore(s *discord.Session, m *discord.MessageCreate) {
	sinbin := quarantineRole(s, m.ChannelID, m.GuildID)
	if sinbin != "" {
		target := m.Mentions[0].ID
		restoreRoles(s, m.GuildID, target)

		err := s.GuildMemberRoleRemove(m.GuildID, target, sinbin)
		sendSuccessOrFail(s, m.ChannelID, err, "remove", target)
	}
}

func quarantineRole(s *discord.Session, channelID string, guildID string) string {
	roles, _ := s.GuildRoles(guildID)
	for _, r := range roles {
		if r.Name == "Quarantine" {
			return r.ID
		}
	}
	s.ChannelMessageSend(channelID, "I couldn't find the **@Quarantine** role!")
	return ""
}

func memberRoles(s *discord.Session, guildID string, target string) []string {
	member, err := s.GuildMember(guildID, target)
	roles := make([]string, 0)
	if err == nil {
		roles = append(roles, member.Roles...)
	}
	return roles
}

func removeRoles(s *discord.Session, g string, u string, roles []string) {
	for _, role := range roles {
		fmt.Printf("Removing role %s...\n", role)
		s.GuildMemberRoleRemove(g, u, role)
	}
}

func restoreRoles(s *discord.Session, g string, u string) {
	roles := make([]string, 0)
	if util.ReadJSON(util.GuildPath("m.roles", g, u), &roles) != nil {
		return
	}
	for _, role := range roles {
		s.GuildMemberRoleAdd(g, u, role)
	}
}

func sendSuccessOrFail(s *discord.Session, channelID string, err error, mode string, target string) {
	op, result := "adding", "now in Quarantine."
	if mode == "remove" {
		op, result = "removing", "back out of Quarantine!"
	}
	if err != nil {
		s.ChannelMessageSend(channelID, "Oops! Couldn't "+op+" the **@Quarantine** role.")
		fmt.Println("\nError:\n", err)
		return
	}
	s.ChannelMessageSend(channelID, fmt.Sprintf("<@%s> is %s", target, result))
}
