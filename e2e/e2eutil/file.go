package e2eutil

import (
	"io/ioutil"
	"testing"
)

// RunCommand loads a file from the given path.
func ReadFile(t *testing.T, path string) string {
	t.Helper()

	b, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
