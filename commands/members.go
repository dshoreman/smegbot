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
	withRole := membersWithRole(members, role.ID)
	count := len(withRole)
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The **@%s** role has **%d** members:", role.Name, count))
	if count == 0 {
		return
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("```\n%s\n```", strings.Join(withRole, "\n")))
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
