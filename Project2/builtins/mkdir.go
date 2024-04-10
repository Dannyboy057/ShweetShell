package builtins

import (
	"fmt"
	"os"
)

// creates a new directory with the specified name.
func MakeDirectory(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing directory name")
	}

	err := os.Mkdir(args[0], 0755)
	if err != nil {
		return err
	}

	fmt.Printf("Directory '%s' created successfully.\n", args[0])
	return nil
}
