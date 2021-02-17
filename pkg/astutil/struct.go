package astutil

import (
	"go/ast"
	"reflect"
)

// GetStructField gets the field name or matching tagged name
func GetStructField(field *ast.Field, tag string) string {
	key, found := reflect.StructTag(field.Tag.Value).Lookup(tag)

	if found {
		return key
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

	return ""
}
