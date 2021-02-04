package parser

import (
	"path/filepath"

	"github.com/go-utils/gopackages"
	"golang.org/x/xerrors"
)

func Parse(dir string) (*Group, error) {

}

type parser struct {
	module    string
	moduleDir string
	rootDir   string
}

func newParser(dir string) (*parser, error) {
	modulePath, err := gopackages.GetGoModPath(dir)

	if err != nil {
		return nil, xerrors.Errorf("failed to find go.mod: %w", err)
	}

	moduleDir := filepath.Dir(modulePath)

	module, err := gopackages.GetGoModule(modulePath)

	if err != nil {
		return nil, xerrors.Errorf("failed to parse go.mod: %w", err)
	}

	return &parser{
		module:    module,
		moduleDir: moduleDir,
		rootDir:   dir,
	}, nil
}

func (p *parser) relativePathFromRoot(dir string) (string, error) {
	d, err := filepath.Rel(p.moduleDir, dir)

	if err != nil {
		return "", xerrors.Errorf("failed to get relative path from module root: %w", err)
	}

	return filepath.Clean(d), nil
}

func (p *parser) walk(dir string) (*Group, error) {

}
