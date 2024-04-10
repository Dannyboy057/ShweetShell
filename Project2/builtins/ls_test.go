package builtins_test

import (
	"os"
	"testing"

	"github.com/Dannyboy057/ShweetShell/Project2/builtins"
)

func TestListDirectoryContents(t *testing.T) {
	t.Parallel()

	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Create some files in the temporary directory
	files := []string{"file1.txt", "file2.txt", "dir1", "dir2"}
	for _, name := range files {
		path := tmpDir + "/" + name
		if name[:3] == "dir" {
			os.Mkdir(path, 0755)
		} else {
			f, _ := os.Create(path)
			f.Close()
		}
	}

	// Test listing all files and directories
	if err := builtins.ListDirectoryContents(); err != nil {
		t.Fatalf("Error listing directory contents: %v", err)
	}
}
