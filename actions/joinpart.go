package actions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	dg "github.com/bwmarrin/discordgo"
)

func onJoin(s *dg.Session, m *dg.GuildMemberAdd) {
	s.ChannelMessageSend(getChannel(s, m.GuildID), fmt.Sprintf("<@%s> has joined! :wave:", m.User.ID))
}

func onChange(s *dg.Session, m *dg.GuildMemberUpdate) {
	saveNick(m.GuildID, m.User, m.Nick)
}

func onPart(s *dg.Session, m *dg.GuildMemberRemove) {
	s.ChannelMessageSend(getChannel(s, m.GuildID), fmt.Sprintf("<@%s> has left :slight_frown:", m.User.ID))
}

func getChannel(s *dg.Session, guildID string) string {
	channels, _ := s.GuildChannels(guildID)

	return channels[1].ID
}

func saveNick(guildID string, u *dg.User, nick string) {
	path := filepath.Join("./storage/guilds", guildID, "members", u.ID)
	err := os.MkdirAll(path, 0700)
	if err != nil {
		fmt.Println("\nError: Couldn't create directory "+path, err)
		return
	}

	err = ioutil.WriteFile(filepath.Join(path, "nick.txt"), []byte(nick), 0644)
	if err != nil {
		fmt.Println("\nError: Couldn't write nick.txt", err)
	} else {
		fmt.Printf("\nWritten new nick for @%s#%s: %s\n", u.Username, u.Discriminator, nick)
	}
}
