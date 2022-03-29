package e2eutil

import (
	"os"
	"testing"
)

// ReadFile loads a file from the given path.
func ReadFile(t *testing.T, path string) string {
	t.Helper()

	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
