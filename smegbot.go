package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	discord "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/actions"
	"github.com/dshoreman/smegbot/cli"
	"github.com/dshoreman/smegbot/commands"
	flag "github.com/ogier/pflag"
)

var (
	token string
)

func init() {
	cli.PrintLogo()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n\n")
		flag.PrintDefaults()
	}

	flag.StringVarP(&token, "token", "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	if token == "" {
		cli.Die("Token must be set.", nil)
	}

	dg, err := discord.New("Bot " + token)
	if err != nil {
		cli.Die("Could not create session.", err)
	}

	err = dg.Open()
	if err != nil {
		cli.Die("Could not connect to Discord.", err)
	}

	actions.Register(dg)
	dg.AddHandler(commands.OnMessageReceived)

	fmt.Print("Connected! Press Ctrl-C to exit.\n\n")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sig

	fmt.Println("\n\nDisconnecting...")
	dg.Close()
	fmt.Println("Goodbye!")
}
