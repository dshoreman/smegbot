package actions

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
)

func onGuildJoin(s *dg.Session, m *dg.GuildCreate) {
	printGuildInfo(s, m.Guild)
}

func printGuildInfo(s *dg.Session, g *dg.Guild) {
	joinDate, _ := g.JoinedAt.Parse()
	f, output := "15:04:05 on 2nd January 2006", `Smegbot has joined %s!
   Server has %d members and is owned by <@%s>.
   Smegbot joined at %s. System messages are sent to <#%s>.
`
	fmt.Printf(output, g.Name, g.MemberCount, g.OwnerID, joinDate.Format(f), g.SystemChannelID)
}
