// Package server generates server-side library
package server

import (
	"os"
	"testing"

	"github.com/go-generalize/api_gen/v2/pkg/parser"
)

func TestGenerate(t *testing.T) {
	gr, err := parser.Parse("../../samples/standard/api")

	if err != nil {
		t.Fatal(err)
	}

	_ = os.RemoveAll("./testdir/output")

	gen, err := NewGenerator(gr, "./testdir/output", "main", "devel")

	if err != nil {
		t.Fatal(err)
	}

	if err := gen.Generate(); err != nil {
		t.Fatal(err)
	}
}

// TestGenerateAliasRepro is a regression test for *types.Alias support (Go 1.23+).
// With GODEBUG=gotypesalias=1 (the default since Go 1.23), type aliases such as
// "any" (alias for interface{}) appear as *types.Alias in go/types output.
// go-easyparser v0.3.3 (used in api_gen v2.14.1) lacked a case for *types.Alias and
// panicked with "unsupported named type: *types.Alias" for any interface that uses
// map[string]any. This test verifies that code generation completes without error.
// The static fixture under testdir/alias_repro/server/ is intentionally NOT
// added to make gen_samples to avoid false-positive diffs from version-string drift.
func TestGenerateAliasRepro(t *testing.T) {
	gr, err := parser.Parse("./testdir/alias_repro/interfaces")

	if err != nil {
		t.Fatalf("parser.Parse failed (regression: *types.Alias support): %v", err)
	}

	// Use a subdirectory under testdir so gopackages.NewModule can find testdir/go.mod.
	outDir := "./testdir/alias_repro/server"
	_ = os.RemoveAll(outDir)

	gen, err := NewGenerator(gr, outDir, "server", "devel")

	if err != nil {
		t.Fatalf("NewGenerator failed: %v", err)
	}

	if err := gen.Generate(); err != nil {
		t.Fatalf("Generate failed (regression: *types.Alias support): %v", err)
	}
}
