package cli

import (
	"fmt"
	"os"
)

// Version is the current Smegbot version
const Version = "1.2.0"

// PrintLogo prints the Smegbot logo with Version info
func PrintLogo() {
	fmt.Printf(`
8""""8                    8""""8
8      eeeeeee eeee eeeee 8    8   eeeee eeeee
8eeeee 8  8  8 8    8   8 8eeee8ee 8   8   8
    88 8  8  8 8ee  8     88     8 8   8   8
e   88 8  8  8 8    8  "8 88     8 8   8   8
8eee88 8  8  8 88ee 88ee8 88eeeee8 8eee8   8
                              Version %s

Initialising...
`, Version)
}

// Die prints an error message and exits the process
func Die(m string, e error) {
	if e == nil {
		fmt.Printf("\nError: %s\nAborting.\n", m)
	} else {
		fmt.Printf("\nError: %s\n  %s\nAborting.\n", m, e)
	}
	os.Exit(1)
}
