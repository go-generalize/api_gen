package server

import (
	"bytes"
	"go/format"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

func (g *Generator) generateAPIErrorPackage(apierrorPath string) error {
	if err := os.MkdirAll(apierrorPath, 0777); err != nil {
		return xerrors.Errorf("failed to mkdir all %s: %w", apierrorPath, err)
	}

	fp, err := os.Create(filepath.Join(apierrorPath, "apierror.go"))
	if err != nil {
		return xerrors.Errorf("failed to create %s: %w", apierrorPath, err)
	}
	defer fp.Close()

	buf := bytes.NewBuffer(nil)

	if err := g.apierrorTemplate.Execute(buf, map[string]string{
		"XerrorsPackage": "golang.org/x/xerrors",
	}); err != nil {
		return xerrors.Errorf("failed to execute template: %w", err)
	}

	src, err := format.Source(buf.Bytes())

	if err != nil {
		return xerrors.Errorf("failed to format code: %w", err)
	}

	io.Copy(fp, bytes.NewBuffer(src))

	return nil
}
