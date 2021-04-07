package server

import (
	"embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/go-utils/gopackages"
	"golang.org/x/xerrors"
)

//go:embed templates/*.tmpl
var templates embed.FS

// Generator generates handlers
type Generator struct {
	group  *parser.Group
	base   string
	module *gopackages.Module

	controllerTemplate        *template.Template
	controllerBundlerTemplate *template.Template
	propsTemplate             *template.Template
	apierrorTemplate          *template.Template

	AppVersion string
}

func NewGenerator(group *parser.Group, base, version string) (*Generator, error) {
	g := &Generator{
		group:      group,
		base:       base,
		AppVersion: version,

		controllerTemplate:        template.Must(template.ParseFS(templates, "templates/controller_template.go.tmpl")),
		controllerBundlerTemplate: template.Must(template.ParseFS(templates, "templates/controller_bundler_template.go.tmpl")),
		propsTemplate:             template.Must(template.ParseFS(templates, "templates/props_template.go.tmpl")),
		apierrorTemplate:          template.Must(template.ParseFS(templates, "templates/apierror_template.go.tmpl")),
	}

	module, err := gopackages.NewModule(base)
	if err != nil {
		return nil, xerrors.Errorf("failed to analyze module for %s: %w", base, err)
	}

	g.module = module

	return g, nil
}

func (g *Generator) Generate() error {
	propsPath := filepath.Join(g.base, "props")

	propsPackage, err := g.module.GetImportPath(propsPath)
	if err != nil {
		return xerrors.Errorf("failed to get import path for %s: %w", propsPath, err)
	}

	if err := g.generateProps(propsPath); err != nil {
		return xerrors.Errorf("failed to generate props in %s: %w", propsPath, err)
	}

	apierrorPath := filepath.Join(g.base, "pkg/apierror")

	apierrorPackage, err := g.module.GetImportPath(apierrorPath)
	if err != nil {
		return xerrors.Errorf("failed to get import path for %s: %w", propsPath, err)
	}

	if err := g.generateAPIErrorPackage(apierrorPath); err != nil {
		return xerrors.Errorf("failed to generate props in %s: %w", apierrorPackage, err)
	}

	endpoints, err := g.generateControllers(filepath.Join(g.base, "controller"), propsPackage, g.group)
	if err != nil {
		return xerrors.Errorf("failed to generate controllers: %w", err)
	}

	err = g.generateControllerBundler(endpoints, propsPackage, apierrorPackage)
	if err != nil {
		return xerrors.Errorf("failed to generate controller bundler: %w", err)
	}

	return nil
}

func (g *Generator) generateControllers(
	root, propsPackage string, gr *parser.Group,
) ([]*bundlerEndpoint, error) {
	if err := os.MkdirAll(root, 0777); err != nil {
		return nil, xerrors.Errorf("failed to mkdir %s: %w", root, err)
	}

	endpoints := make([]*bundlerEndpoint, 0, len(gr.Endpoints))
	for _, ep := range gr.Endpoints {
		be, err := g.generateController(root, propsPackage, ep)

		if err != nil {
			return nil, xerrors.Errorf("failed to generate controller for %s: %w", ep.GetFullPath("/", func(rawPath, path, placeholder string) string {
				return path
			}), err)
		}

		endpoints = append(endpoints, be)
	}

	for _, child := range gr.Children {
		be, err := g.generateControllers(root, propsPackage, child)

		if err != nil {
			return nil, xerrors.Errorf("failed to generate %s: %w", child.Path, err)
		}

		endpoints = append(endpoints, be...)
	}

	return endpoints, nil
}
