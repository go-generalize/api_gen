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

	"github.com/go-generalize/api_gen/v2/pkg/parser"
	"github.com/go-generalize/api_gen/v2/pkg/util"
	"github.com/iancoleman/strcase"
	"golang.org/x/xerrors"
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
	Name            string
	Method          string
	Path            string
	Request         string
	Response        string
	MultipartUpload bool
	FileFields      []parser.FileField
}

type group struct {
	ShortName, Name string
	Children        []*group
	Endpoints       []*endpoint
}

type generator struct {
	PackageName     string
	Groups          []*group
	imports         map[string]string
	Imports         []importPair
	Version         string
	Root            *group
	MultipartUpload bool

	baseGroup          *parser.Group
	typesDirImportPath string
}

type importPair struct {
	Alias, Path string
}

// GenerateClient generates a Go client for api_gen
func (g *generator) GenerateClient() (string, error) {
	var err error

	g.imports = make(map[string]string)
	g.Root, err = g.generateGroup(g.baseGroup)

	if err != nil {
		return "", xerrors.Errorf("failed to convert groups/endpoints: %w", err)
	}

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

	tmpl, err := template.New("client.go.tmpl").Funcs(
		template.FuncMap{
			"IsArrayMultipleUpload": func(f parser.FileField) bool {
				return f.Type == parser.UploadMultipleFiles
			},
		},
	).ParseFS(clientGoTemplate, "templates/client.go.tmpl")

	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	if err := tmpl.ExecuteTemplate(buf, "client.go.tmpl", g); err != nil {
		return "", err
	}

	formatted, err := format.Source(buf.Bytes())

	if err != nil {
		return "", err
	}

	return string(formatted), nil
}

func (g *generator) generateEndpoint(ep *parser.Endpoint) (*endpoint, error) {
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

	if ep.UseMultipartUpload {
		g.MultipartUpload = true
		gen.MultipartUpload = true

		fields, err := parser.GetFileFields(ep.RequestGo2tsPayload)

		if err != nil {
			return nil, xerrors.Errorf("failed to get file fields: %w", err)
		}

		gen.FileFields = fields
	}

	fullPath = `"` + fullPath + `"`

	gen.Path = fullPath
	gen.Name = strcase.ToCamel(strings.ToLower(string(ep.Method))) + ep.RawPath

	importAlias := g.imports[ep.GetParent().ImportPath]

	gen.Request = fmt.Sprintf("%s.%s", importAlias, ep.RequestPayloadName)
	gen.Response = fmt.Sprintf("%s.%s", importAlias, ep.ResponsePayloadName)

	return gen, nil
}

func (g *generator) generateGroup(gr *parser.Group) (*group, error) {
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
		gr, err := g.generateGroup(gr.Children[i])

		if err != nil {
			return nil, xerrors.Errorf("failed to generate group(%s): %w", gr.Children[i].Name, err)
		}

		gen.Children = append(gen.Children, gr)
	}

	gen.Endpoints = make([]*endpoint, 0, len(gr.Endpoints))
	for i := range gr.Endpoints {
		ep, err := g.generateEndpoint(gr.Endpoints[i])

		if err != nil {
			return nil, xerrors.Errorf("failed to generate endpoint(%s): %w", gr.Endpoints[i].RawPath, err)
		}

		gen.Endpoints = append(gen.Endpoints, ep)
	}

	g.Groups = append(g.Groups, gen)

	return gen, nil
}
