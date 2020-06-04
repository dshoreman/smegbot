package commands

import (
	"fmt"
	"regexp"
	"strings"

	dg "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/config"
)

func listConfigValues(s *dg.Session, m *dg.MessageCreate) {
	config.LoadGuild(m.GuildID)

	joins, parts := config.Guild.JoinChannel, config.Guild.PartChannel
	if joins == "" {
		joins = "Not set"
	}
	if parts == "" {
		parts = "Not set"
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
		"> \n> **Current configuration**:\n> \n> `joinChannel`  -  <#%s>\n> `partChannel`  -  <#%s>\n",
		joins, parts))
}

func setConfigOption(s *dg.Session, m *dg.MessageCreate) {
	config.LoadGuild(m.GuildID)

	save, msg := false, "Nothing to save! Did you actually change anything? :person_facepalming:"
	matched, _ := regexp.MatchString(`\.config (join|part)Channel <#[0-9]+>`, m.Content)
	if matched {
		args := strings.Fields(m.Content)
		if args[1] == "joinChannel" && config.Guild.JoinChannel != args[2] {
			config.Guild.JoinChannel = args[2][2 : len(args[2])-1]
			save = true
		} else if args[1] == "partChannel" && config.Guild.PartChannel != args[2] {
			config.Guild.PartChannel = args[2][2 : len(args[2])-1]
			save = true
		}
	} else {
		msg = "That's not a valid config command :slight_frown:"
	}

	if save {
		msg = "Config saved! :tada:"
		err := config.SaveGuild(m.GuildID)
		if err != nil {
			msg = "I couldn't save the config :sob:"
			fmt.Println("Error: Failed saving guild config\n", err)
		}
	}
	s.ChannelMessageSend(m.ChannelID, msg)
}
