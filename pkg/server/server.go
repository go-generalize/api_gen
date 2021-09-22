// Package server generates server-side library
package server

import (
	"embed"
	"path/filepath"
	"text/template"

	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/go-utils/gopackages"
	"golang.org/x/sync/singleflight"
	"golang.org/x/xerrors"
)

//go:embed templates/*.tmpl
var templates embed.FS

// Generator generates handlers
type Generator struct {
	group             *parser.Group
	base, basePackage string
	module            *gopackages.Module

	controllerTemplate            *template.Template
	controllerBundlerTemplate     *template.Template
	controllerInitializerTemplate *template.Template
	propsTemplate                 *template.Template
	apierrorTemplate              *template.Template
	mockTemplate                  *template.Template
	mockControllerTemplate        *template.Template

	AppVersion string

	sfGroup singleflight.Group
}

// NewGenerator returns a new Generator
func NewGenerator(group *parser.Group, base, basePackage, version string) (*Generator, error) {
	g := &Generator{
		group:       group,
		base:        base,
		basePackage: basePackage,
		AppVersion:  version,

		controllerTemplate:            template.Must(template.ParseFS(templates, "templates/controller_template.go.tmpl")),
		controllerBundlerTemplate:     template.Must(template.ParseFS(templates, "templates/controller_bundler_template.go.tmpl")),
		controllerInitializerTemplate: template.Must(template.ParseFS(templates, "templates/controller_initializer_template.go.tmpl")),
		propsTemplate:                 template.Must(template.ParseFS(templates, "templates/props_template.go.tmpl")),
		apierrorTemplate:              template.Must(template.ParseFS(templates, "templates/apierror_template.go.tmpl")),
		mockTemplate:                  template.Must(template.ParseFS(templates, "templates/mock_template.go.tmpl")),
		mockControllerTemplate:        template.Must(template.ParseFS(templates, "templates/mock_controller_template.go.tmpl")),
	}

	module, err := gopackages.NewModule(base)
	if err != nil {
		return nil, xerrors.Errorf("failed to analyze module for %s: %w", base, err)
	}

	g.module = module

	return g, nil
}

// Generate generates controllers and related packages in the specified directory
func (g *Generator) Generate() error {
	propsPath := filepath.Join(g.base, "props")

	propsPackage, err := g.module.GetImportPath(propsPath)
	if err != nil {
		return xerrors.Errorf("failed to get import path for %s: %w", propsPath, err)
	}

	apierrorPath := filepath.Join(g.base, "pkg/apierror")

	apierrorPackage, err := g.module.GetImportPath(apierrorPath)
	if err != nil {
		return xerrors.Errorf("failed to get import path for %s: %w", propsPath, err)
	}

	if err := g.generateProps(propsPath); err != nil {
		return xerrors.Errorf("failed to generate props in %s: %w", propsPath, err)
	}

	err = g.generateControllerInitializer(g.base, propsPackage, apierrorPackage)
	if err != nil {
		return xerrors.Errorf("failed to generate controller initializer: %w", err)
	}

	if err := g.generateAPIErrorPackage(apierrorPath); err != nil {
		return xerrors.Errorf("failed to generate props in %s: %w", apierrorPackage, err)
	}

	controllerPath := filepath.Join(g.base, "controller")
	endpoints, err := g.generateControllers(
		controllerPath,
		g.group,
		func(ep *parser.Endpoint) (*bundlerEndpoint, error) {
			return g.generateController(controllerPath, propsPackage, ep)
		},
	)
	if err != nil {
		return xerrors.Errorf("failed to generate controllers: %w", err)
	}

	if err = g.generateControllerBundler(
		endpoints,
		propsPackage,
		apierrorPackage,
		"bundler.go",
		false,
	); err != nil {
		return xerrors.Errorf("failed to generate controller bundler: %w", err)
	}

	if err := g.generateMock(apierrorPackage, propsPackage); err != nil {
		return xerrors.Errorf("failed to generate mock: %w", err)
	}

	return nil
}
