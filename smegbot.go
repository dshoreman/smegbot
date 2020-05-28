package main

import (
	"fmt"
	"os"

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
		fmt.Println("Error: Token must be set. Aborting.")
		os.Exit(1)
	}
}
