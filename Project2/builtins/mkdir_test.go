package builtins_test

import (
	"os"
	"testing"

	"github.com/Dannyboy057/ShweetShell/Project2/builtins"
)

func TestMakeDirectory(t *testing.T) {
	t.Parallel()

	// Temporary directory for testing
	tmpDir := t.TempDir()

	// Directory name for testing
	dirName := "testdir"

	// Test making a new directory
	err := builtins.MakeDirectory(tmpDir + "/" + dirName)
	if err != nil {
		t.Fatalf("Error making directory: %v", err)
	}

	// Check if directory was created
	_, err = os.Stat(tmpDir + "/" + dirName)
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("Directory '%s' was not created", dirName)
		}
		t.Fatalf("Error checking directory: %v", err)
	}
}
