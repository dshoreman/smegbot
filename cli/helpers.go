package cli

import (
	"fmt"
	"os"
)

// PrintLogo prints the Smegbot logo with Version info
func PrintLogo(version string) {
	fmt.Printf(`
8""""8                    8""""8
8      eeeeeee eeee eeeee 8    8   eeeee eeeee
8eeeee 8  8  8 8    8   8 8eeee8ee 8   8   8
    88 8  8  8 8ee  8     88     8 8   8   8
e   88 8  8  8 8    8  "8 88     8 8   8   8
8eee88 8  8  8 88ee 88ee8 88eeeee8 8eee8   8
                              Version %s

Initialising...
`, version)
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
