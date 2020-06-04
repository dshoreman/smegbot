package util

import (
	"path/filepath"
)

const basePath string = "./storage/guilds"

// GuildPath returns paths to files in a guild's config dir
func GuildPath(parts ...string) string {
	g, f1, u, f2 := parts[1], "members", "", ""
	switch parts[0] {
	case "m.nick":
		u, f2 = parts[2], "nick.txt"
	case "m.roles":
		u, f2 = parts[2], "roles.txt"
	case "member":
		u = parts[2]
	case "config":
		f1 = "config.json"
	default:
		f1 = ""
	}
	return filepath.Join(basePath, g, f1, u, f2)
}
