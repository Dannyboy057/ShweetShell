package builtins

import (
	"fmt"
)

// help displays the available commands
func ShowHelp() error {
	helpMessage := `
	   Available commands:
	   - cd [directory]: Change the current directory.
	   - env: Display environment variables.
	   - exit: Exit the shell.
	   - ls [options: List files and/or directories in the current directory.
	   - clear: Clear the console.
	   - pwd: Display the current working directory.
	   - mkdir [directory]: Create a new directory.
	   - rm: takes file's or directory's name as argument and removes it.
	   - help: Display this help message.
`
	fmt.Println(helpMessage)
	return nil
}
