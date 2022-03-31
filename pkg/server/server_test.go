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
