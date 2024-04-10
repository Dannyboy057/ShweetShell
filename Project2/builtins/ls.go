package builtins

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// lists files and directories in the current directory
func ListDirectoryContents(args ...string) error {
	// Ensure we only received one argument
	if len(args) == 0 {
		args = append(args, "all") // default to listing all contents if no arguments are provided
	}

	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Read the directory contents
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// Print the directory contents based on the argument
	switch strings.ToLower(args[0]) {
	case "all":
		for _, file := range files {
			fmt.Println(file.Name())
		}
	case "d":
		for _, file := range files {
			if file.IsDir() {
				fmt.Println(file.Name())
			}
		}
	case "f":
		for _, file := range files {
			if !file.IsDir() {
				fmt.Println(file.Name())
			}
		}
	default:
		return fmt.Errorf("unsupported argument: %s", args[0])
	}

	return nil
}
