package server

import (
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

func (g *Generator) generateProps(propsPath string) error {
	if err := os.MkdirAll(propsPath, 0777); err != nil {
		return xerrors.Errorf("failed to mkdir all %s: %w", propsPath, err)
	}

	fp, err := os.Create(filepath.Join(propsPath, "props.go"))
	if err != nil {
		return xerrors.Errorf("failed to create %s: %w", propsPath, err)
	}
	defer fp.Close()

	if err := g.propsTemplate.Execute(fp, nil); err != nil {
		return xerrors.Errorf("failed to execute template: %w", err)
	}

	return nil
}
