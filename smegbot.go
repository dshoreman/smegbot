package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	discord "github.com/bwmarrin/discordgo"
	flag "github.com/ogier/pflag"
)

var (
	token string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n\n")
		flag.PrintDefaults()
	}

	flag.StringVarP(&token, "token", "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	fmt.Println("8\"\"\"\"8                    8\"\"\"\"8               ")
	fmt.Println("8      eeeeeee eeee eeeee 8    8   eeeee eeeee ")
	fmt.Println("8eeeee 8  8  8 8    8   8 8eeee8ee 8   8   8   ")
	fmt.Println("    88 8  8  8 8ee  8     88     8 8   8   8   ")
	fmt.Println("e   88 8  8  8 8    8  \"8 88     8 8   8   8   ")
	fmt.Println("8eee88 8  8  8 88ee 88ee8 88eeeee8 8eee8   8   ")

	fmt.Println("\nInitialising...")

	if token == "" {
		fmt.Println("\nError: Token must be set. Aborting.")
		os.Exit(1)
	}

	dg, err := discord.New("Bot " + token)
	if err != nil {
		fmt.Println("\nError: Could not create session.\n", err)
		os.Exit(1)
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("\nError: Could not connect to Discord\n", err)
		os.Exit(1)
	}

	dg.AddHandler(onMessageReceived)

	fmt.Print("Connected! Press Ctrl-C to exit.\n\n")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sig

	fmt.Println("\n\nDisconnecting...")
	dg.Close()
	fmt.Println("Goodbye!")
}

func onMessageReceived(s *discord.Session, m *discord.MessageCreate) {
	fmt.Println(m.Author.Username, ":", m.Content)

	// Don't process our own messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		return
	}

	if len(m.Content) > 7 && m.Content[0:7] == ".roles " {
		user := m.Mentions[0]

		member, err := s.GuildMember(m.GuildID, user.ID)
		if err != nil {
			fmt.Println("\nError: Could not get guild member\n", err)
			s.ChannelMessageSend(m.ChannelID, "Are you sure they're still a member?")
			return
		}

		if len(member.Roles) == 0 {
			s.ChannelMessageSend(m.ChannelID, "This user has no roles!")
			return
		}

		roles := make([]string, len(member.Roles))
		for i, roleID := range member.Roles {
			role, err := s.State.Role(m.GuildID, roleID)
			if err != nil {
				fmt.Println("\nError: Could not get role\n", err)
				roles[i] = "•" + roleID
				continue
			}
			roles[i] = fmt.Sprintf("• %s - %s", role.ID, role.Name)
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(
			"**@%s#%s has the following roles:**\n%s",
			user.Username, user.Discriminator,
			strings.Join(roles, "\n"),
		))
		return
	}
}
