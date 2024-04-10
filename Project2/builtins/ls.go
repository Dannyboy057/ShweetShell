package builtins

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// ListDirectoryContents lists files and directories in the current directory based on the argument.
func ListDirectoryContents(args ...string) error {
	// If no arguments provided, default to "all"
	if len(args) == 0 {
		args = append(args, "all")
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

	// Print the column names
	fmt.Println("Mode       LastWriteTime	Length  	Name")

	// Print the directory contents based on the argument
	switch args[0] {
	case "all":
		for _, file := range files {
			info, err := formatFileInfo(file)
			if err != nil {
				return err
			}
			fmt.Println(info)
		}
	case "d":
		for _, file := range files {
			if file.IsDir() {
				info, err := formatFileInfo(file)
				if err != nil {
					return err
				}
				fmt.Println(info)
			}
		}
	case "f":
		for _, file := range files {
			if !file.IsDir() {
				info, err := formatFileInfo(file)
				if err != nil {
					return err
				}
				fmt.Println(info)
			}
		}
	default:
		return fmt.Errorf("unsupported argument: %s", args[0])
	}

	return nil
}

// formatFileInfo formats file information including mode, last write time, and length.
func formatFileInfo(file os.FileInfo) (string, error) {
	mode := file.Mode().String()
	modTime := file.ModTime().Format("Jan _2 15:04")
	size := strconv.FormatInt(file.Size(), 10)
	name := file.Name()

	return fmt.Sprintf("%s %s %12s %20s", mode, modTime, size, name), nil
}
