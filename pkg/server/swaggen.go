// Package server generates server-side library
package server

import (
	"bytes"
	"fmt"
	gotypes "go/types"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/go-generalize/api_gen/v2/pkg/parser"
	"github.com/go-generalize/go-easyparser/types"
)

func (g *Generator) generateSwagComment(ep *parser.Endpoint) string {
	obj := ep.RequestGo2tsPayload

	placeholders := map[string]struct{}{}

	path := ep.GetFullPath("/", func(rawPath, path, placeholder string) string {
		if placeholder != "" {
			placeholders[placeholder] = struct{}{}

			return "{" + placeholder + "}"
		}

		return path
	})

	params := make([]string, 0)
	for k, v := range obj.Entries {
		t, attrs := getSwagType(v.Type)

		param, ok := reflect.StructTag(v.RawTag).Lookup("param")
		if !ok {
			param = k
		}

		if attrs != "" {
			attrs = " " + attrs
		}

		if _, ok := placeholders[param]; ok {
			params = append(params, fmt.Sprintf(`// @Param %s path %s %v "%s"%s`+"\n", param, t, true, k, attrs))
		} else if ep.Method == parser.GET {
			example, ok := reflect.StructTag(v.RawTag).Lookup("example")
			if !ok {
				example = k
			}
			params = append(params, fmt.Sprintf(`// @Param %s query %s %v "%s"%s`+"\n", k, t, !v.Optional, example, attrs))
		}
	}

	if ep.Method != parser.GET {
		params = append(params, fmt.Sprintf(`// @Param body body types.%s true "request parameter"`+"\n", ep.RequestPayloadName))
	}

	sort.Strings(params)

	buf := bytes.NewBuffer(nil)

	if len(ep.SwagComments) != 0 {
		buf.WriteString(strings.Join(ep.SwagComments, "\n") + "\n")
	}
	buf.WriteString("// @Accept json\n")
	buf.WriteString("// @Produce json\n")
	buf.WriteString(strings.Join(params, ""))

	buf.WriteString(fmt.Sprintf("// @Success 200 {object} types.%s\n", ep.ResponsePayloadName))
	buf.WriteString(fmt.Sprintf("// @Router %s [%s]\n", path, strings.ToLower(string(ep.Method))))

	return buf.String()
}

const (
	swagParamInteger = "integer"
	swagParamString  = "string"
	swagParamBoolean = "boolean"
)

func getSwagType(t types.Type) (typeName string, attrs string) {
	switch t := t.(type) {
	case *types.Nullable:
		return getSwagType(t.Inner)
	case *types.Number:
		enumArray := make([]string, len(t.Enum))
		for i, e := range t.Enum {
			enumArray[i] = strconv.FormatInt(e, 10)
		}
		enums := strings.Join(enumArray, ", ")

		if len(enums) != 0 {
			enums = "Enums(" + enums + ")"
		}

		if t.RawType == gotypes.Float32 || t.RawType == gotypes.Float64 {
			return "number", enums
		}

		return swagParamInteger, enums
	case *types.Boolean:
		return swagParamBoolean, ""
	case *types.String:
		enums := strings.Join(t.Enum, ", ")
		if len(enums) != 0 {
			enums = "Enums(" + enums + ")"
		}

		return swagParamString, enums
	case *types.Array:
		typeName, _ = getSwagType(t.Inner)
		return "[]" + typeName, ""
	case *types.Date:
		return swagParamString, "Format(date-time)"
	default:
		return "interface{}", ""
	}
}
