package actions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	dg "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/config"
)

func onChange(s *dg.Session, m *dg.GuildMemberUpdate) {
	g, u := m.GuildID, m.User.ID

	if !nickIsCached(g, u) || currentNick(g, u) != m.Nick {
		saveNick(g, m.User, m.Nick)
	}
}

func onJoin(s *dg.Session, m *dg.GuildMemberAdd) {
	nick, action := currentNick(m.GuildID, m.User.ID), "joined"
	if nick != "" {
		action = "returned"
	}
	sendJoinPart(s, m.GuildID, m.User, nick, action)
}

func onPart(s *dg.Session, m *dg.GuildMemberRemove) {
	nick := currentNick(m.GuildID, m.User.ID)
	sendJoinPart(s, m.GuildID, m.User, nick, "left")
}

func sendJoinPart(s *dg.Session, g string, u *dg.User, nick string, action string) {
	emoji, nickstring := "wave", ""
	if action == "left" {
		emoji = "slight_frown"
	}
	if nick != "" && nick != u.Username {
		nickstring = "\nYou may know them as *" + nick + "*."
	}
	s.ChannelMessageSend(getChannel(s, g, action), fmt.Sprintf("**@%s#%s** has %s! :%s:%s",
		u.Username, u.Discriminator, action, emoji, nickstring))
}

func getChannel(s *dg.Session, guildID string, action string) string {
	config.LoadGuild(guildID)
	j, p := config.Guild.JoinChannel, config.Guild.PartChannel

	if action == "left" && p != "" {
		return p
	} else if action != "left" && j != "" {
		return j
	}
	channels, _ := s.GuildChannels(guildID)
	return channels[1].ID
}

func nickIsCached(g string, u string) bool {
	path := filepath.Join("./storage/guilds/", g, "members", u, "nick.txt")
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

func currentNick(g string, u string) string {
	path := filepath.Join("./storage/guilds", g, "members", u, "nick.txt")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Couldn't read "+path, err)
		return ""
	}
	return string(b)
}

func saveNick(guildID string, u *dg.User, nick string) {
	if nick == "" {
		fmt.Println("\nNick was removed, saving username instead.")
		nick = u.Username
	}

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
