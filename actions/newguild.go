package actions

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/util"
)

func onGuildJoin(s *dg.Session, m *dg.GuildCreate) {
	printGuildInfo(s, m.Guild)
}

func printGuildInfo(s *dg.Session, g *dg.Guild) {
	joinDate, _ := g.JoinedAt.Parse()
	f, output := "15:04:05 on January 2, 2006", `Smegbot has joined %s! %s
   Smegbot first joined at %s. Of %d members, %d have had names saved.
   Guild is owned by <@%s>. System messages are sent to <#%s>.

`
	fmt.Printf(output, g.Name, hasConfig(g),
		joinDate.Format(f), g.MemberCount, withNames(g),
		g.OwnerID, g.SystemChannelID)
}

func hasConfig(g *dg.Guild) string {
	if util.FileExists(util.GuildPath("config", g.ID)) {
		return "It appears to have config already."
	}
	return "Seems it's missing its config too!"
}

func withNames(g *dg.Guild) int {
	count := 0
	for _, m := range g.Members {
		f := util.GuildPath("m.nick", g.ID, m.User.ID)
		if util.FileExists(f) && util.ReadString(f) != "" {
			count++
		}
	}
	return count
}
