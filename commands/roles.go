package commands

import (
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

func listMemberRoles(s *dg.Session, m *dg.MessageCreate) {
	g, u := m.GuildID, m.Mentions[0]
	gm, err := s.GuildMember(g, u.ID)
	if err != nil {
		fmt.Println("\nError: Could not get guild member\n", err)
		s.ChannelMessageSend(m.ChannelID, "Are you sure they're still a member?")
		return
	}
	if len(gm.Roles) == 0 {
		s.ChannelMessageSend(m.ChannelID, "This user has no roles!")
		return
	}

	roles := make([]string, len(gm.Roles))
	for i, r := range gm.Roles {
		role, err := s.State.Role(g, r)
		if err != nil {
			fmt.Println("\nError: Could not get role\n", err)
			roles[i] = "• " + r
			continue
		}
		roles[i] = fmt.Sprintf("• @%s", role.Name)
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("**@%s#%s** has the following **%d** roles:\n```\n%s\n```",
		u.Username, u.Discriminator, len(roles), strings.Join(roles, "\n")))
}
