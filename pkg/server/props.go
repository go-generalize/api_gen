package server

import (
	"bytes"
	"go/format"
	"io"
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

	buf := bytes.NewBuffer(nil)
	if err := g.propsTemplate.Execute(fp, nil); err != nil {
		return xerrors.Errorf("failed to execute template: %w", err)
	}

	src, err := format.Source(buf.Bytes())

	if err != nil {
		return xerrors.Errorf("failed to format code: %w", err)
	}

	if _, err := io.Copy(fp, bytes.NewBuffer(src)); err != nil {
		return xerrors.Errorf("failed to write props.go: %w", err)
	}

	return nil
}
