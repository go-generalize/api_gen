package server

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/go-generalize/api_gen/pkg/parser"
)

func (g *Generator) generateSwagComment(ep *parser.Endpoint) string {
	buf := bytes.NewBuffer(nil)

	buf.WriteString("// @Accept  json\n")
	buf.WriteString("// @Produce  json\n")

	if ep.Method == parser.GET {
		// TODO: Implement
	} else {
		buf.WriteString(fmt.Sprintf(`// @Param body body types.%s true "request parameter"`+"\n", ep.RequestPayloadName))
	}

	path := ep.GetFullPath("/", func(rawPath, path, placeholder string) string {
		if placeholder != "" {
			buf.WriteString(fmt.Sprintf(`// @Param %s path %s true "%s"`+"\n", placeholder, "string", placeholder))

			return "{" + placeholder + "}"
		}

		return path
	})

	buf.WriteString(fmt.Sprintf("// @Success 200 {object} types.%s\n", ep.ResponsePayloadName))
	buf.WriteString(fmt.Sprintf("// @Success 200 {object} types.%s\n", ep.ResponsePayloadName))
	buf.WriteString(fmt.Sprintf("// @Router %s [%s]\n", path, strings.ToLower(string(ep.Method))))

	return buf.String()
}
