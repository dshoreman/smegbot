package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

// Config stores configuration for a guild
type Config struct {
	JoinChannel string `json:"channels.join"`
	PartChannel string `json:"channels.part"`
}

var config Config

func listConfigValues(s *dg.Session, m *dg.MessageCreate) {
	b, err := ioutil.ReadFile(filepath.Join("./storage/guilds", m.GuildID, "config.json"))
	if err == nil {
		json.Unmarshal(b, &config)
	}

	joins, parts := config.JoinChannel, config.PartChannel
	if joins == "" {
		joins = "Not set"
	}
	if parts == "" {
		parts = "Not set"
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
		"> \n> **Current configuration**:\n> \n> `joinChannel`  -  %s\n> `partChannel`  -  %s\n",
		joins, parts))
}
