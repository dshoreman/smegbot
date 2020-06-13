package util

import (
	"path/filepath"

	dg "github.com/bwmarrin/discordgo"
)

const basePath string = "./storage/guilds"

// GuildMember attemps to grab a Member from state, falling back to API if necessary
func GuildMember(s *dg.Session, g string, u string) (*dg.Member, error) {
	m, err := s.State.Member(g, u)
	if err != nil {
		m, err = s.GuildMember(g, u)
	}
	return m, err
}

// IsAdmin checks if a member has Administrator permission on any role
func IsAdmin(s *dg.Session, g string, u string) (bool, error) {
	m, err := GuildMember(s, g, u)
	if err != nil {
		return false, err
	}
	for _, roleID := range m.Roles {
		r, err := s.State.Role(g, roleID)
		if err != nil {
			return false, err
		}
		if r.Permissions&dg.PermissionAdministrator != 0 {
			return true, nil
		}
	}
	return false, nil
}

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
