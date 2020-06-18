package actions

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/config"
	"github.com/dshoreman/smegbot/util"
)

func onGuildJoin(s *dg.Session, m *dg.GuildCreate) {
	g := m.Guild
	needsConfig, current := !hasConfig(g.ID), savedNameCount(g)

	printGuildInfo(g, current)
	if needsConfig || current < g.MemberCount {
		fmt.Printf("Running setup for %s:\n", g.Name)
	}
	if needsConfig {
		writeConfig(g)
	}
	if current < len(g.Members) {
		saveMemberNames(g, current)
	}
}

func printGuildInfo(g *dg.Guild, nameCount int) {
	joinDate, _ := g.JoinedAt.Parse()
	f, output := "15:04:05 on January 2, 2006", `Smegbot has joined %s!
   Smegbot first joined at %s. Of %d members, %d have had names saved.
   Guild is owned by <@%s>. System messages are sent to <#%s>.

`
	fmt.Printf(output, g.Name, joinDate.Format(f), g.MemberCount, nameCount, g.OwnerID, g.SystemChannelID)
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

func saveMemberNames(g *dg.Guild, nameCount int) {
	saved, missing := 0, g.MemberCount-nameCount
	fmt.Printf("* Updating member names (missing %d)... ", missing)

	for _, m := range g.Members {
		if _, err := util.SaveMemberName(g.ID, m); err == nil {
			saved++
		}
	}
	fmt.Printf("%d/%d OK\n", saved, g.MemberCount)
}

func hasConfig(g string) bool {
	return util.FileExists(util.GuildPath("config", g))
}

func savedNameCount(g *dg.Guild) int {
	count := 0
	for _, m := range g.Members {
		f := util.GuildPath("m.nick", g.ID, m.User.ID)
		if util.FileExists(f) && util.ReadString(f) != "" {
			count++
		}
	}
	return count
}
