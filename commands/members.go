package commands

import (
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

func listRoleMembers(s *dg.Session, m *dg.MessageCreate) {
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
	msg := "None of the members seem to have the **@" + role.Name + "** role. :slight_frown:"
	withRole := membersWithRole(members, role.ID)

	if len(withRole) > 0 {
		msg = fmt.Sprintf("\nThere are **%d** member(s) with the **@%s** role:\n```\n%s\n```",
			len(withRole), role.Name, strings.Join(withRole, "\n"))
	}
	s.ChannelMessageSend(m.ChannelID, msg)
}

func memberHasRole(member *dg.Member, role string) bool {
	for _, current := range member.Roles {
		if current == role {
			return true
		}
	}
	return false
}

func membersWithRole(members []*dg.Member, roleID string) []string {
	m := make([]string, 0)
	for _, member := range members {
		if !memberHasRole(member, roleID) {
			continue
		}

		u, nick := member.User, ""
		if member.Nick != "" {
			nick = "- " + member.Nick
		}

		m = append(m, fmt.Sprintf("• @%s#%s %s", u.Username, u.Discriminator, nick))
	}
	return m
}
