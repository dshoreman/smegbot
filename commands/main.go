package commands

import (
	"fmt"

	discord "github.com/bwmarrin/discordgo"
)

// OnMessageReceived processes incoming messages from Discord to register commands
func OnMessageReceived(s *discord.Session, m *discord.MessageCreate) {
	fmt.Println(m.Author.Username, ":", m.Content)

	if m.Author.ID != s.State.User.ID {
		runAll(s, m)
	}
}

func runAll(s *discord.Session, m *discord.MessageCreate) {
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
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
