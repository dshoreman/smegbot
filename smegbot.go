package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	discord "github.com/bwmarrin/discordgo"
	"github.com/dshoreman/smegbot/commands"
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

	dg.AddHandler(commands.OnMessageReceived)

	fmt.Print("Connected! Press Ctrl-C to exit.\n\n")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sig

	fmt.Println("\n\nDisconnecting...")
	dg.Close()
	fmt.Println("Goodbye!")
}
