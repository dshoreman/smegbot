package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GuildConfig defines valid guild config options
type GuildConfig struct {
	JoinChannel string `json:"channels.join"`
	PartChannel string `json:"channels.part"`
}

// Guild stores a guild's configuration
var Guild GuildConfig

// LoadGuild loads the config for a guild
func LoadGuild(guildID string) {
	b, err := ioutil.ReadFile(filepath.Join("./storage/guilds", guildID, "config.json"))
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
	path := filepath.Join("./storage/guilds", guildID)
	err = os.MkdirAll(path, 0700)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(path, "config.json"), b, 0644)
}
