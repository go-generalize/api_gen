// Package clientts generates client-side TypeScript library
package clientts

import (
	"fmt"
	"path/filepath"

	"github.com/go-generalize/api_gen/v2/pkg/parser"
	"github.com/go-generalize/go-easyparser/types"
	go2tsgenerator "github.com/go-generalize/go2ts/pkg/generator"
	go2tstypes "github.com/go-generalize/go2ts/pkg/types"
	"golang.org/x/xerrors"
)

// GenerateTypes generates request/response types in TypeScript
func (g *generator) GenerateTypes(fn func(relPath, code string) error) error {
	err := g.generateTypes(g.root, fn)

	if err != nil {
		return xerrors.Errorf("failed to generate request/response types: %w", err)
	}

	return nil
}

const fileHeaderName = "mime/multipart.FileHeader"

func (g *generator) generateTypes(gr *parser.Group, fn func(relPath, code string) error) error {
	for _, child := range gr.Children {
		if err := g.generateTypes(child, fn); err != nil {
			return xerrors.Errorf("an error occurred in %s: %w", gr.Dir, err)
		}
	}

	if len(gr.Endpoints) == 0 {
		return nil
	}

	replaceFileHeaderAll(gr.ParsedTypes)

	gen := go2tsgenerator.NewGenerator(gr.ParsedTypes)

	gen.CustomGenerator = func(t go2tstypes.Type) (generated string, union bool) {
		o, ok := t.(*go2tstypes.Object)

		if !ok {
			return "", false
		}

		if o.Name == fileHeaderName {
			return "File | Blob", true
		}

		return "", false
	}

	code := gen.Generate()

	relative := gr.GetFullPath(string(filepath.Separator), func(rawPath, path, placeholder string) string {
		return rawPath
	})

	code = fmt.Sprintf(headerComment, g.AppVersion) + code

	p := filepath.Join("classes", relative, "types.ts")
	if err := fn(p, code); err != nil {
		return xerrors.Errorf("failed to save %s: %w", p, err)
	}

	return nil
}

func replaceFileHeaderAll(t map[string]types.Type) {
	for _, v := range t {
		v, ok := v.(*types.Object)

		if !ok {
			continue
		}

		replaceFileHeader(v)
	}
}

func replaceFileHeader(obj *types.Object) {
	for k, v := range obj.Entries {
		if obj, ok := v.Type.(*types.Object); ok {
			replaceFileHeader(obj)

			continue
		}

		t, err := parser.ValidateMultipartUploadType(v.Type, v.RawTag)

		if err != nil || t == parser.UploadNone {
			continue
		}

		if t == parser.UploadSingleFile {
			v.Type = &types.Object{
				Name: fileHeaderName,
			}

		} else {
			v.Type =
				&types.Array{
					Inner: &types.Object{
						Name: fileHeaderName,
					},
				}

		}
		v.Optional = true

		obj.Entries[k] = v
	}
}
