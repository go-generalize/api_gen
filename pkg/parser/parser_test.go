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

	pp.SetColorScheme(pp.ColorScheme{
		Bool:            pp.NoColor,
		Integer:         pp.NoColor,
		Float:           pp.NoColor,
		String:          pp.NoColor,
		StringQuotation: pp.NoColor,
		EscapedChar:     pp.NoColor,
		FieldName:       pp.NoColor,
		PointerAdress:   pp.NoColor,
		Nil:             pp.NoColor,
		Time:            pp.NoColor,
		StructName:      pp.NoColor,
		ObjectLength:    pp.NoColor,
	})

	fmt.Println(pp.Sprint(gr))
}
