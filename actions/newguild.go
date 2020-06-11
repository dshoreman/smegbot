package actions

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
)

func onGuildJoin(s *dg.Session, m *dg.GuildCreate) {
	fmt.Println("Smegbot has joined", m.Guild.Name)
}
