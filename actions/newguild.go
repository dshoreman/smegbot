package actions

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/config"
	"github.com/dshoreman/smegbot/util"
)

func onGuildJoin(s *dg.Session, m *dg.GuildCreate) {
	printGuildInfo(s, m.Guild)

	if !hasConfig(m.Guild.ID) || withNames(m.Guild) < len(m.Guild.Members) {
		fmt.Printf("Running setup for %s:\n", m.Guild.Name)
	}
	if !hasConfig(m.Guild.ID) {
		writeConfig(m.Guild)
	}
	if withNames(m.Guild) < len(m.Guild.Members) {
		saveMemberNames(m.Guild)
	}
}

func printGuildInfo(s *dg.Session, g *dg.Guild) {
	joinDate, _ := g.JoinedAt.Parse()
	f, output := "15:04:05 on January 2, 2006", `Smegbot has joined %s! %s
   Smegbot first joined at %s. Of %d members, %d have had names saved.
   Guild is owned by <@%s>. System messages are sent to <#%s>.

`
	configured := "Seems it's missing its config too!"
	if hasConfig(g.ID) {
		configured = "It appears to have config already."
	}
	fmt.Printf(output, g.Name, configured,
		joinDate.Format(f), g.MemberCount, withNames(g),
		g.OwnerID, g.SystemChannelID)
}

func hasConfig(g string) bool {
	return util.FileExists(util.GuildPath("config", g))
}

func writeConfig(g *dg.Guild) {
	fmt.Printf("* Saving initial config... ")
	config.Guild.JoinChannel = g.SystemChannelID
	config.Guild.PartChannel = g.SystemChannelID

	if err := config.SaveGuild(g.ID); err != nil {
		fmt.Println("ERROR\n   ", err)
		return
	}
	fmt.Println("OK")
}

func saveMemberNames(g *dg.Guild) {
	saved, missing := 0, g.MemberCount-withNames(g)
	fmt.Printf("* Updating member names (missing %d)... ", missing)

	for _, m := range g.Members {
		n, f := m.Nick, util.GuildPath("m.nick", g.ID, m.User.ID)
		if n == "" {
			n = m.User.String()
		}
		if err := util.WriteFile(f, []byte(n)); err == nil {
			saved++
		}
	}
	fmt.Printf("%d/%d OK\n", saved, g.MemberCount)
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
