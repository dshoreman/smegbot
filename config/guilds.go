package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/dshoreman/smegbot/util"
)

// GuildConfig defines valid guild config options
type GuildConfig struct {
	JoinChannel string `json:"channels.join"`
	PartChannel string `json:"channels.part"`
}

// Guild stores a guild's configuration
var Guild GuildConfig

// LoadGuild loads the config for a guild
func LoadGuild(g string) {
	b, err := ioutil.ReadFile(util.GuildPath("config", g))
	if err == nil {
		json.Unmarshal(b, &Guild)
	}
}

// SaveGuild saves the config for a guild
func SaveGuild(guildID string) error {
	b, err := json.Marshal(Guild)
	if err != nil {
		return err
	}
	err = os.MkdirAll(util.GuildPath("", guildID), 0700)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(util.GuildPath("config", guildID), b, 0644)
}
