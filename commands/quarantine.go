package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	discord "github.com/bwmarrin/discordgo"
)

func nuke(s *discord.Session, m *discord.MessageCreate) {
	sinbin := quarantineRole(s, m.ChannelID, m.GuildID)
	if sinbin != "" {
		target := m.Mentions[0].ID
		roles := memberRoles(s, m.GuildID, target)

		saved, err := backupRoles(s, m.ChannelID, m.GuildID, target, roles)
		if err != nil {
			fmt.Println("\nError: Could not save roles to file\n", err)
		}
		err = s.GuildMemberRoleAdd(m.GuildID, target, sinbin)
		if saved && err == nil {
			removeRoles(s, m.GuildID, target, roles)
		}
		sendSuccessOrFail(s, m.ChannelID, err, "add", target)
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

func backupRoles(s *discord.Session, channelID string, g string, u string, roles []string) (bool, error) {
	b, err := json.Marshal(roles)
	if err != nil {
		return false, err
	}
	err = os.MkdirAll(util.GuildPath("member", g, u), 0700)
	if err == nil {
		err = ioutil.WriteFile(util.GuildPath("m.roles", g, u), b, 0644)
	}
	return err == nil, err
}

func removeRoles(s *discord.Session, g string, u string, roles []string) {
	for _, role := range roles {
		fmt.Printf("Removing role %s...\n", role)
		s.GuildMemberRoleRemove(g, u, role)
	}
}

func restoreRoles(s *discord.Session, g string, u string) {
	b, err := ioutil.ReadFile(util.GuildPath("m.roles", g, u))
	if err != nil {
		fmt.Println("Couldn't read user's roles.json", err)
		return
	}
	roles := make([]string, 0)
	json.Unmarshal(b, &roles)
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
