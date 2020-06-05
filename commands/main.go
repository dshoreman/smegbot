package commands

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/cli"
)

// OnMessageReceived processes incoming messages from Discord to register commands
func OnMessageReceived(s *dg.Session, m *dg.MessageCreate) {
	fmt.Println(m.Author.Username, ":", m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	ok, err := hasPermission(s, m)
	if err != nil {
		fmt.Println("Failed permissions check.", err)
	}
	if ok {
		runAll(s, m)
	}
}

func runAll(s *dg.Session, m *dg.MessageCreate) {
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		return
	}
	if m.Content == ".version" {
		s.ChannelMessageSend(m.ChannelID, "Currently running Smegbot version `"+cli.Version+"`")
	}

	if m.Content == ".config" {
		listConfigValues(s, m)
		return
	}
	if len(m.Content) > 8 && m.Content[0:8] == ".config " {
		setConfigOption(s, m)
		return
	}

	if len(m.Content) > 9 && m.Content[0:9] == ".members " {
		listRoleMembers(s, m)
		return
	}
	if len(m.Content) > 7 && m.Content[0:7] == ".roles " {
		listMemberRoles(s, m)
		return
	}

	if len(m.Content) > 6 && m.Content[0:6] == ".nuke " {
		nuke(s, m)
		return
	}
	if len(m.Content) > 9 && m.Content[0:9] == ".restore " {
		restore(s, m)
		return
	}
}

func hasPermission(s *dg.Session, m *dg.MessageCreate) (bool, error) {
	guildID, userID := m.GuildID, m.Author.ID
	member, err := s.State.Member(guildID, userID)
	if err != nil {
		if member, err = s.GuildMember(guildID, userID); err != nil {
			return false, err
		}
	}
	for _, roleID := range member.Roles {
		role, err := s.State.Role(guildID, roleID)
		if err != nil {
			return false, err
		}
		if role.Permissions&dg.PermissionAdministrator != 0 {
			return true, nil
		}
	}
	return false, nil
}
