package actions

import (
	discord "github.com/bwmarrin/discordgo"
)

// Register sets up all non-message event handlers
func Register(dg *discord.Session) {
	dg.AddHandler(func(s *discord.Session, m *discord.GuildMemberAdd) {
		onJoin(s, m)
	})
	dg.AddHandler(func(s *discord.Session, m *discord.GuildMemberRemove) {
		onPart(s, m)
	})
}
