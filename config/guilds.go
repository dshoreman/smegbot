package config

import (
	"github.com/dshoreman/smegbot/util"
)

// GuildConfig defines valid guild config options
type GuildConfig struct {
	SuperUser string `json:"admin.user"`
	AdminRole string `json:"admin.role"`

	JoinChannel string `json:"channels.join"`
	PartChannel string `json:"channels.part"`

	QuarantineRole string `json:"roles.quarantine"`
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
