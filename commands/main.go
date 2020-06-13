package commands

import (
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/cli"
	"github.com/dshoreman/smegbot/util"
)

// OnMessageReceived processes incoming messages from Discord to register commands
func OnMessageReceived(s *dg.Session, m *dg.MessageCreate) {
	fmt.Println(m.Author.Username, ":", m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}
	if ok, _ := util.IsAdmin(s, m.GuildID, m.Author.ID); ok {
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
		return
	}

	if m.Content == ".config" {
		listConfigValues(s, m)
		return
	}
	if strings.HasPrefix(m.Content, ".config ") {
		setConfigOption(s, m)
		return
	}

	if strings.HasPrefix(m.Content, ".members ") {
		listRoleMembers(s, m)
		return
	}
	if strings.HasPrefix(m.Content, ".roles ") {
		listMemberRoles(s, m)
		return
	}

	if strings.HasPrefix(m.Content, ".nuke ") {
		nuke(s, m)
		return
	}
	if strings.HasPrefix(m.Content, ".restore ") {
		restore(s, m)
		return
	}
}
