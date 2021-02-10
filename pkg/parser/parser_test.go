package parser

import (
	"fmt"
	"testing"

	"github.com/k0kubun/pp"
)

func TestParse(t *testing.T) {
	gr, err := Parse("../../server_generator/sample")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(pp.Sprint(gr))
}
