// Package clientgo generates client-side Go library
package clientgo

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/go-generalize/go2go"
	go2tsparser "github.com/go-generalize/go2ts/pkg/parser"
	"golang.org/x/xerrors"
)

const headerComment = `// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: %s

`

func parserFilter(opt *go2tsparser.FilterOpt) bool {
	if opt.Dependency {
		return true
	}
	if !opt.BasePackage {
		return false
	}
	if !opt.Exported {
		return false
	}

	return strings.HasSuffix(opt.Name, "Request") || strings.HasSuffix(opt.Name, "Response")
}

func (g *generator) GenerateTypes(fn func(relPath, code string) error) error {
	err := g.generateTypes(g.baseGroup, fn)

	if err != nil {
		return xerrors.Errorf("failed to generate request/response types: %w", err)
	}

	return nil
}

func (g *generator) generateTypes(gr *parser.Group, fn func(relPath, code string) error) error {
	for _, child := range gr.Children {
		if err := g.generateTypes(child, fn); err != nil {
			return xerrors.Errorf("an error occurred in %s: %w", gr.Dir, err)
		}
	}

	if len(gr.Endpoints) == 0 {
		return nil
	}

	psr, err := go2tsparser.NewParser(gr.Dir, parserFilter)

	if err != nil {
		return xerrors.Errorf("failed to parse %s: %w", gr.Dir, err)
	}

	parsed, err := psr.Parse()

	if err != nil {
		return xerrors.Errorf("failed to parse %s: %w", gr.Dir, err)
	}

	endpointStructs := make([]string, 0, 20)
	for _, ep := range gr.Endpoints {
		endpointStructs = append(endpointStructs, gr.ImportPath+"."+ep.RequestPayloadName)
		endpointStructs = append(endpointStructs, gr.ImportPath+"."+ep.ResponsePayloadName)
	}

	code, err := go2go.NewGenerator(parsed, endpointStructs).Generate()

	relative := gr.GetFullPath(string(filepath.Separator), func(rawPath, path, placeholder string) string {
		return rawPath
	})

	code = fmt.Sprintf(headerComment, g.Version) + code

	p := filepath.Join("classes", relative, "types.go")
	if err := fn(p, code); err != nil {
		return xerrors.Errorf("failed to save %s: %w", p, err)
	}

	return nil
}
