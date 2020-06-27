package commands

import (
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/cli"
	"github.com/dshoreman/smegbot/config"
	"github.com/dshoreman/smegbot/util"
)

// OnMessageReceived processes incoming messages from Discord to register commands
func OnMessageReceived(s *dg.Session, m *dg.MessageCreate) {
	if m.GuildID == "" {
		fmt.Printf("[DM] @%s: %s\n", m.Author.String(), m.Content)
	} else {
		g, _ := s.State.Guild(m.GuildID)
		c, _ := s.State.Channel(m.ChannelID)
		fmt.Printf("[%s] @%s in #%s: %s\n", g.Name, m.Author.String(), c.Name, m.Content)
	}

	if m.Author.ID == s.State.User.ID {
		return
	}
	config.LoadGuild(m.GuildID)

	if ok, _ := util.IsAdmin(s, util.MemberCheck{
		Guild:  m.GuildID,
		Member: m.Author.ID,
		Root:   config.Guild.SuperUser,
	}); ok {
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
