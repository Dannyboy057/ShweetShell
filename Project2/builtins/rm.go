package builtins

import (
	"fmt"
	"os"
	"path/filepath"
)

// Remove removes files or directories with the specified names if they exist in the current directory.
func Remove(names ...string) error {
	for _, name := range names {
		path := filepath.Join(".", name)

		// Check if the file or directory exists
		_, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				// File or directory does not exist, continue to the next one
				continue
			}
			// Other error occurred
			return fmt.Errorf("failed to check existence of %s: %w", name, err)
		}

		// Remove the file or directory
		err = os.RemoveAll(path)
		if err != nil {
			return fmt.Errorf("failed to remove %s: %w", name, err)
		}
	}
	return nil
}
