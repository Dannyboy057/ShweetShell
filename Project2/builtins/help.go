package builtins

import (
	"fmt"
)

// help displays the available commands
func ShowHelp() error {
	helpMessage := `
	   Available commands:
- cd <directory>: Change directory.
- env: Print environment variables.
- exit: Exit the shell.
- pwd: Print the current working directory.
- clear: Clear the console screen.
- help: Display this help message.
`
	fmt.Println(helpMessage)
	return nil
}
