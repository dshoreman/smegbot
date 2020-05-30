package commands

import (
	"fmt"
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

func listRoleMembers(s *discord.Session, m *discord.MessageCreate) {
	role, err := s.State.Role(m.GuildID, m.MentionRoles[0])
	if err != nil {
		fmt.Println("\nError: Could not get role\n", err)
		return
	}

	members, err := s.GuildMembers(m.GuildID, "", 1000)
	if err != nil {
		fmt.Println("\nError: Failed loading guild members\n", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Searching %d members...", len(members)))

	withRole := make([]string, 0)
	for _, member := range members {
		if !memberHasRole(member, role.ID) {
			continue
		}

		nick := ""
		if member.Nick != "" {
			nick = "\n  -- " + member.Nick
		}

		withRole = append(withRole, fmt.Sprintf("• %s: @%s#%s %s",
			member.User.ID, member.User.Username, member.User.Discriminator, nick,
		))
	}

	if len(withRole) > 0 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
			"\nThere are **%d** member(s) with the **@%s** role:\n```\n%s\n```",
			len(withRole), role.Name, strings.Join(withRole, "\n"),
		))
		return
	}

	s.ChannelMessageSend(m.ChannelID,
		"None of the members seem to have the **@"+role.Name+"** role. :slight_frown:")
	return
}

func memberHasRole(member *discord.Member, role string) bool {
	for _, current := range member.Roles {
		if current == role {
			return true
		}
	}
	return false
}
