package actions

import (
	dg "github.com/bwmarrin/discordgo"
	cmd "github.com/dshoreman/smegbot/commands"
)

// Register sets up all non-message event handlers
func Register(s *dg.Session) {
	s.AddHandler(onGuildJoin)

	s.AddHandler(onJoin)
	s.AddHandler(onChange)
	s.AddHandler(onPart)

	s.AddHandler(cmd.OnMessageReceived)
}
