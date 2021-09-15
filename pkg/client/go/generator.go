// Package clientgo generates client-side Go library
package clientgo

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"sort"
	"strings"
	"text/template"

	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/go-generalize/api_gen/pkg/util"
	"github.com/iancoleman/strcase"
)

//go:embed templates/client.go.tmpl
var clientGoTemplate embed.FS

// Generator generates a Go client for api_gen
type Generator interface {
	GenerateClient() (string, error)
	GenerateTypes(fn func(relPath, code string) error) error
}

var _ Generator = &generator{}

// NewGenerator generates a Go client for api_gen
func NewGenerator(gr *parser.Group, typesDirImportPath, packageName, version string) Generator {
	gen := &generator{
		PackageName:        packageName,
		Version:            version,
		typesDirImportPath: typesDirImportPath,
		baseGroup:          gr,
	}

	return gen
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
	imports     map[string]string
	Imports     []importPair
	Version     string
	Root        *group

	baseGroup          *parser.Group
	typesDirImportPath string
}

type importPair struct {
	Alias, Path string
}

// GenerateClient generates a Go client for api_gen
func (g *generator) GenerateClient() (string, error) {
	g.imports = make(map[string]string)
	g.Root = g.generateGroup(g.baseGroup)

	for k, v := range g.imports {
		// Overwrite import path to ./classes
		path := g.typesDirImportPath + strings.TrimPrefix(k, g.baseGroup.ImportPath)

		g.Imports = append(g.Imports, importPair{
			Alias: v,
			Path:  path,
		})
	}

	sort.Slice(g.Imports, func(i, j int) bool {
		return g.Imports[i].Path+g.Imports[i].Path < g.Imports[j].Path+g.Imports[j].Path
	})

	tmpl, err := template.ParseFS(clientGoTemplate, "templates/client.go.tmpl")

	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, g); err != nil {
		return "", err
	}

	formatted, err := format.Source(buf.Bytes())

	if err != nil {
		return "", err
	}

	return string(formatted), nil
}

func (g *generator) generateEndpoint(ep *parser.Endpoint) *endpoint {
	gen := &endpoint{}

	gen.Method = string(ep.Method)

	fullPath := ep.GetFullPath("/", func(rawPath, path, placeholder string) string {
		if placeholder != "" {
			return fmt.Sprintf(
				`" + fmt.Sprint(reqPayload.%s) + "`,
				util.FindStructField(ep.RequestGo2tsPayload, parser.QueryParamTag, placeholder),
			)
		}

		return path
	})
	fullPath = `"` + fullPath + `"`

	gen.Path = fullPath
	gen.Name = strcase.ToCamel(strings.ToLower(string(ep.Method))) + ep.RawPath

	importAlias := g.imports[ep.GetParent().ImportPath]

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

	if len(gr.Endpoints) != 0 {
		importAlias := gr.GetFullPath("_", func(rawPath, path, placeholder string) string {
			return rawPath
		})
		if importAlias == "" {
			importAlias = "root"
		}

		g.imports[gr.ImportPath] = importAlias
	}

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
