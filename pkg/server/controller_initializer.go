// Package server generates server-side library
package server

import (
	"bytes"
	"go/format"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

func (g *Generator) generateControllerInitializer(initializerPath, propsPackage string) error {
	fp, err := os.Create(filepath.Join(initializerPath, "controller_initializer.go"))
	if err != nil {
		return xerrors.Errorf("failed to create %s: %w", initializerPath, err)
	}
	defer fp.Close()

	buf := bytes.NewBuffer(nil)
	if err := g.controllerInitializerTemplate.Execute(fp, map[string]string{
		"ControllerPropsPackage": propsPackage,
		"Package":                g.basePackage,
		"AppVersion":             g.AppVersion,
	}); err != nil {
		return xerrors.Errorf("failed to execute template: %w", err)
	}

	src, err := format.Source(buf.Bytes())

	if err != nil {
		return xerrors.Errorf("failed to format code: %w", err)
	}

	if _, err := io.Copy(fp, bytes.NewBuffer(src)); err != nil {
		return xerrors.Errorf("failed to write controller_initializer.go: %w", err)
	}

	return nil
}
