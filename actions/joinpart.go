package actions

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/config"
	"github.com/dshoreman/smegbot/util"
)

func onChange(s *dg.Session, m *dg.GuildMemberUpdate) {
	g, u := m.GuildID, m.User.ID

	if !util.FileExists(util.GuildPath("m.nick", g, u)) || currentNick(g, u) != m.Nick {
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

func currentNick(g string, u string) string {
	return util.ReadString(util.GuildPath("m.nick", g, u))
}

func saveNick(g string, u *dg.User, nick string) {
	if nick == "" {
		fmt.Println("\nNick was removed, saving username instead.")
		nick = u.Username
	}

	err := util.WriteFile(util.GuildPath("m.nick", g, u.ID), []byte(nick))
	if err != nil {
		fmt.Println("\nError: Couldn't write nick.txt", err)
		return
	}
	fmt.Printf("\nWritten new nick for @%s#%s: %s\n", u.Username, u.Discriminator, nick)
}
