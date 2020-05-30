package commands

import (
	"fmt"

	discord "github.com/bwmarrin/discordgo"
)

func nuke(s *discord.Session, m *discord.MessageCreate) {
	sinbin := ""
	guildRoles, _ := s.GuildRoles(m.GuildID)
	for _, role := range guildRoles {
		if role.Name == "Quarantine" {
			sinbin = role.ID
		}
	}
	if sinbin == "" {
		s.ChannelMessageSend(m.ChannelID, "I couldn't find the **@Quarantine** role!")
		return
	}

	target := m.Mentions[0].ID
	err := s.GuildMemberRoleAdd(m.GuildID, target, sinbin)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Something went wrong trying to add the **@Quarantine** role.")
		fmt.Println("\nError: Could not add Quarantine role.\n", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> is now in Quarantine.", target))
	return
}

func restore(s *discord.Session, m *discord.MessageCreate) {
	sinbin := ""
	guildRoles, _ := s.GuildRoles(m.GuildID)
	for _, role := range guildRoles {
		if role.Name == "Quarantine" {
			sinbin = role.ID
		}
	}
	if sinbin == "" {
		s.ChannelMessageSend(m.ChannelID, "I couldn't find the **@Quarantine** role!")
		return
	}

	target := m.Mentions[0].ID
	err := s.GuildMemberRoleRemove(m.GuildID, target, sinbin)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Something went wrong removing the **@Quarantine** role.")
		fmt.Println("\nError: Could not remove Quarantine role.\n", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> is back out of Quarantine!", target))
	return
}
