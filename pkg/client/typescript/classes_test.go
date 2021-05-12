package clientts

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/google/go-cmp/cmp"
)

func Test_generator_GenerateTypes(t *testing.T) {
	group, err := parser.Parse("../../../samples/standard/api")

	if err != nil {
		t.Fatal(err)
	}

	g := NewGenerator(group, "2.0")

	err = g.GenerateTypes(func(relPath, code string) error {
		file, err := os.ReadFile(filepath.Join("./testdata/", relPath))

		if err != nil {
			t.Errorf("failed to parse %s: %+v", relPath, err)

			return nil
		}

		if diff := cmp.Diff(string(file), code); diff != "" {
			t.Errorf("failed to parse %s: %s", relPath, diff)
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
