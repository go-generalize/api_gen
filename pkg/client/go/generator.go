package clientgo

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"strings"
	"text/template"

	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/go-utils/astutil"
	"github.com/iancoleman/strcase"
)

const (
	queryTagKey = "param"
)

//go:embed templates/client.go.tmpl
var clientGoTemplate embed.FS

// Generate generates a Go client for api_gen
func Generate(gr *parser.Group, packageName string) (string, error) {
	params := generator{
		PackageName: packageName,
	}

	params.generate(gr)

	tmpl, err := template.ParseFS(clientGoTemplate, "templates/client.go.tmpl")

	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, params); err != nil {
		return "", err
	}

	formatted, err := format.Source(buf.Bytes())

	if err != nil {
		return "", err
	}

	return string(formatted), nil
}

type endpoint struct {
	Name     string
	Method   string
	Path     string
	Request  string
	Response string
}

type group struct {
	ShortName, Name string
	Children        []*group
	Endpoints       []*endpoint
}

type generator struct {
	PackageName string
	Groups      []*group
	Imports     map[string]string
	Root        *group
}

func (g *generator) generate(gr *parser.Group) {
	g.Imports = make(map[string]string)
	g.Root = g.generateGroup(gr)
}

func (g *generator) generateEndpoint(ep *parser.Endpoint) *endpoint {
	gen := &endpoint{}

	gen.Method = string(ep.Method)

	fullPath := ep.GetFullPath("/", func(rawPath, path, placeholder string) string {
		if placeholder != "" {
			return fmt.Sprintf(
				`" + fmt.Sprint(reqPayload.%s) + "`,
				astutil.FindStructField(ep.RequestPayload, queryTagKey, placeholder),
			)
		}

		return path
	})
	fullPath = `"` + fullPath + `"`

	gen.Path = fullPath
	gen.Name = strcase.ToCamel(strings.ToLower(string(ep.Method))) + ep.RawPath

	importAlias := g.Imports[ep.GetParent().ImportPath]

	gen.Request = fmt.Sprintf("%s.%s", importAlias, ep.RequestPayloadName)
	gen.Response = fmt.Sprintf("%s.%s", importAlias, ep.ResponsePayloadName)

	return gen
}

func (g *generator) generateGroup(gr *parser.Group) *group {
	gen := &group{}

	p := gr.GetFullPath("_", func(rawPath, path, placeholder string) string {
		return rawPath
	})

	gen.ShortName = strcase.ToCamel(gr.RawPath)
	gen.Name = "Group" + p

	importAlias := gr.GetFullPath("_", func(rawPath, path, placeholder string) string {
		return rawPath
	})
	if importAlias == "" {
		importAlias = "root"
	}

	g.Imports[gr.ImportPath] = importAlias

	gen.Children = make([]*group, 0, len(gr.Children))
	for i := range gr.Children {
		gen.Children = append(gen.Children, g.generateGroup(gr.Children[i]))
	}

	gen.Endpoints = make([]*endpoint, 0, len(gr.Endpoints))
	for i := range gr.Endpoints {
		gen.Endpoints = append(gen.Endpoints, g.generateEndpoint(gr.Endpoints[i]))
	}

	g.Groups = append(g.Groups, gen)

	return gen
}
