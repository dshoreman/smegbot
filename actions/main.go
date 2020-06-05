package actions

import (
	dg "github.com/bwmarrin/discordgo"
)

// Register sets up all non-message event handlers
func Register(s *dg.Session) {
	s.AddHandler(func(s *dg.Session, m *dg.GuildMemberAdd) {
		onJoin(s, m)
	})
	s.AddHandler(func(s *dg.Session, m *dg.GuildMemberUpdate) {
		onChange(s, m)
	})
	s.AddHandler(func(s *dg.Session, m *dg.GuildMemberRemove) {
		onPart(s, m)
	})
}
