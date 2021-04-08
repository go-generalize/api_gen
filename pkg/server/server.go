// Package server generates server-side library
package server

import (
	"embed"
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

	controllerTemplate            *template.Template
	controllerBundlerTemplate     *template.Template
	controllerInitializerTemplate *template.Template
	propsTemplate                 *template.Template
	apierrorTemplate              *template.Template

	AppVersion string
}

// NewGenerator returns a new Generator
func NewGenerator(group *parser.Group, base, version string) (*Generator, error) {
	g := &Generator{
		group:      group,
		base:       base,
		AppVersion: version,

		controllerTemplate:            template.Must(template.ParseFS(templates, "templates/controller_template.go.tmpl")),
		controllerBundlerTemplate:     template.Must(template.ParseFS(templates, "templates/controller_bundler_template.go.tmpl")),
		controllerInitializerTemplate: template.Must(template.ParseFS(templates, "templates/controller_initializer_template.go.tmpl")),
		propsTemplate:                 template.Must(template.ParseFS(templates, "templates/props_template.go.tmpl")),
		apierrorTemplate:              template.Must(template.ParseFS(templates, "templates/apierror_template.go.tmpl")),
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

	if err := g.generateProps(propsPath); err != nil {
		return xerrors.Errorf("failed to generate props in %s: %w", propsPath, err)
	}

	err = g.generateControllerInitializer(g.base, propsPackage)
	if err != nil {
		return xerrors.Errorf("failed to generate controller initializer: %w", err)
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
