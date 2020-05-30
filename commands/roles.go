package commands

import (
	"fmt"
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

func listMemberRoles(s *discord.Session, m *discord.MessageCreate) {
	user := m.Mentions[0]

	member, err := s.GuildMember(m.GuildID, user.ID)
	if err != nil {
		fmt.Println("\nError: Could not get guild member\n", err)
		s.ChannelMessageSend(m.ChannelID, "Are you sure they're still a member?")
		return
	}

	if len(member.Roles) == 0 {
		s.ChannelMessageSend(m.ChannelID, "This user has no roles!")
		return
	}

	roles := make([]string, len(member.Roles))
	for i, roleID := range member.Roles {
		role, err := s.State.Role(m.GuildID, roleID)
		if err != nil {
			fmt.Println("\nError: Could not get role\n", err)
			roles[i] = "• " + roleID
			continue
		}
		roles[i] = fmt.Sprintf("• %s: @%s", role.ID, role.Name)
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
		"**@%s#%s** has the following **%d** roles:\n```\n%s\n```",
		user.Username, user.Discriminator,
		len(roles), strings.Join(roles, "\n"),
	))
	return
}
