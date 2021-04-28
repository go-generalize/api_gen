package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	gr, err := Parse("../../samples/standard/api")

	if err != nil {
		t.Fatal(err)
	}

	t.Error(gr)
}
