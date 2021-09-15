package util

import (
	"reflect"

	go2tstypes "github.com/go-generalize/go2ts/pkg/types"
)

// GetStructField gets the field name or matching tagged name
func GetStructField(e *go2tstypes.ObjectEntry, tag string) string {
	if e.RawTag != "" {
		key, found := reflect.StructTag(e.RawTag).Lookup(tag)

		if found {
			return key
		}
	}

	return e.RawName
}

// FindStructField finds the field for struct matching the name
func FindStructField(strct *go2tstypes.Object, tag, name string) string {
	for _, field := range strct.Entries {
		key := GetStructField(&field, tag)

		if key == name {
			return field.RawName
		}
	}

	return ""
}
