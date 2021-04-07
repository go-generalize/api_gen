// Package server generates server-side library
package server

import (
	"os"
	"testing"

	"github.com/go-generalize/api_gen/pkg/parser"
)

func TestGenerate(t *testing.T) {

	gr, err := parser.Parse("../../samples/standard/api")

	if err != nil {
		t.Fatal(err)
	}

	os.RemoveAll("./testfiles")

	gen, err := NewGenerator(gr, "./testfiles", "devel")

	if err != nil {
		t.Fatal(err)
	}

	if err := gen.Generate(); err != nil {
		t.Fatal(err)
	}

}
