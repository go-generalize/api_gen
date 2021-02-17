package astutil

import (
	"go/ast"
	"reflect"

	"github.com/k0kubun/pp"
)

// GetStructField gets the field name or matching tagged name
func GetStructField(field *ast.Field, tag string) string {
	if field.Tag != nil {
		key, found := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1]).Lookup(tag)

		if found {
			return key
		}
	}

	return field.Names[0].Name
}

// FindStructField finds the field for struct matching the name
func FindStructField(strct *ast.StructType, tag, name string) string {
	for _, field := range strct.Fields.List {
		key := GetStructField(field, tag)

		if key == name {
			return field.Names[0].Name
		}
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
	pp.Println(strct, tag, name)

	return ""
}
