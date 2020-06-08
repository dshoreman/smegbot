package config

import (
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
	util.ReadJSON(util.GuildPath("config", g), &Guild)
}

// SaveGuild saves the config for a guild
func SaveGuild(g string) error {
	return util.WriteJSON(util.GuildPath("config", g), Guild)
}
