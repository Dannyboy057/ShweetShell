package builtins

import (
	"fmt"
)

// Clear the console screen.
func ClearConsole() error {
	// ANSI escape sequence
	clearScreen := "\033[2J\033[H"

	// Write the escape sequence to standard output
	_, err := fmt.Print(clearScreen)
	return err
}
