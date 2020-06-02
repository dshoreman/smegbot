package actions

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
)

func onJoin(s *dg.Session, m *dg.GuildMemberAdd) {
	s.ChannelMessageSend(getChannel(s, m.GuildID), fmt.Sprintf("<@%s> has joined! :wave:", m.User.ID))
}

func onPart(s *dg.Session, m *dg.GuildMemberRemove) {
	s.ChannelMessageSend(getChannel(s, m.GuildID), fmt.Sprintf("<@%s> has left :slight_frown:", m.User.ID))
}

func getChannel(s *dg.Session, guildID string) string {
	channels, _ := s.GuildChannels(guildID)

	return channels[1].ID
}