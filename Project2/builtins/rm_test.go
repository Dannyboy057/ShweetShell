package builtins_test

import (
	"os"
	"testing"

	"github.com/Dannyboy057/ShweetShell/Project2/builtins"
)

func TestRemove(t *testing.T) {
	// Create a temporary file for testing
	fileName := "test_file.txt"
	_, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer func() {
		err := os.Remove(fileName)
		if err != nil {
			t.Fatalf("Failed to remove test file: %v", err)
		}
	}()

	// Removing the temporary file
	err = builtins.Remove(fileName)
	if err != nil {
		t.Fatalf("Failed to remove file: %v", err)
	}

	// Check if the file still exists
	_, err = os.Stat(fileName)
	if err == nil {
		t.Fatalf("File %s still exists after removal", fileName)
	}
	if !os.IsNotExist(err) {
		t.Fatalf("Unexpected error when checking file existence: %v", err)
	}

	// Removing a non-existent file
	err = builtins.Remove("nonexistent_file.txt")
	if err != nil {
		t.Fatalf("Failed to remove non-existent file: %v", err)
	}
}
