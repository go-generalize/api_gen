// Package clientdart generates client-side Dart library
package clientdart

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-generalize/api_gen/v2/pkg/parser"
	types "github.com/go-generalize/go-easyparser/types"
	util "github.com/go-generalize/go-easyparser/util"
	go2dartgenerator "github.com/go-generalize/go2dart"
	"golang.org/x/xerrors"
)

// GenerateTypes generates request/response types in Dart
func (g *generator) GenerateTypes(fn func(relPath, code string) error) error {
	err := g.generateTypes(g.root, fn)

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

	gen := go2dartgenerator.NewGenerator(gr.ParsedTypes, nil)

	gen.ExternalImporter = func(o *types.Object) *go2dartgenerator.ExternalImporter {
		rel, err := filepath.Rel(g.root.Dir, o.Position.Filename)

		if err != nil {
			return nil
		}

		if strings.HasPrefix(rel, "../") {
			return nil
		}

		rel, err = filepath.Rel(gr.Dir, filepath.Dir(o.Position.Filename))

		if err != nil {
			return nil
		}

		if rel == "." {
			return nil
		}

		_, structName := util.SplitPackegeStruct(o.Name)

		return &go2dartgenerator.ExternalImporter{
			Path: filepath.Join(rel, "types.dart"),
			Name: structName,
		}
	}

	code, err := gen.Generate()

	if err != nil {
		return xerrors.Errorf("failed to generate: %w", err)
	}

	relative := gr.GetFullPath(string(filepath.Separator), func(rawPath, path, placeholder string) string {
		return rawPath
	})

	code = fmt.Sprintf(headerComment, g.AppVersion) + code

	p := filepath.Join("classes", relative, "types.dart")
	if err := fn(p, code); err != nil {
		return xerrors.Errorf("failed to save %s: %w", p, err)
	}

	return nil
}
