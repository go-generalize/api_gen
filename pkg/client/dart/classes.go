// Package clientdart generates client-side Dart library
package clientdart

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fatih/structtag"
	"github.com/go-generalize/api_gen/v2/pkg/parser"
	"github.com/go-generalize/go-easyparser"
	types "github.com/go-generalize/go-easyparser/types"
	go2dartgenerator "github.com/go-generalize/go2dart"
	"golang.org/x/xerrors"
)

// GenerateTypes generates request/response types in Dart
func (g *generator) GenerateTypes(fn func(relPath, code string) error) error {
	code, err := go2dartgenerator.NewGenerator(nil, nil).Generate()

	if err != nil {
		return xerrors.Errorf("failed to generate common JsonConverter: %w", err)
	}

	if err := fn("common.dart", code); err != nil {
		return xerrors.Errorf("failed to save common.dart: %w", err)
	}

	err = g.generateTypes(g.root, fn)

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

	replaceFileHeaderAll(gr.ParsedTypes)

	gen := go2dartgenerator.NewGenerator(gr.ParsedTypes, nil)

	gen.ExternalImporter = func(o *types.Object) *go2dartgenerator.ExternalImporter {
		if o.Name == fileHeaderName {
			return &go2dartgenerator.ExternalImporter{
				Path: "package:http/http.dart",
				Name: "MultipartFile",
			}
		}

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

		_, structName := easyparser.SplitPackageStruct(o.Name)

		return &go2dartgenerator.ExternalImporter{
			Path: filepath.Join(rel, "types.dart"),
			Name: structName,
		}
	}

	relative := gr.GetFullPath(string(filepath.Separator), func(rawPath, path, placeholder string) string {
		return rawPath
	})

	reversedRelative, err := filepath.Rel(filepath.Join(".", relative), ".")

	if err != nil {
		return xerrors.Errorf("failed to get reversed relative path: %w", err)
	}

	gen.ExternalCommonConverterPath = filepath.Join(reversedRelative, "../common.dart")

	code, err := gen.Generate()

	if err != nil {
		return xerrors.Errorf("failed to generate: %w", err)
	}

	code = fmt.Sprintf(headerComment, g.AppVersion) + code

	p := filepath.Join("classes", relative, "types.dart")
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
		if o, ok := v.Type.(*types.Object); ok {
			replaceFileHeader(o)

			continue
		}

		tags, err := structtag.Parse(v.RawTag)

		if err != nil {
			continue
		}

		// nolint:errcheck
		jsonTag, _ := tags.Get("json")

		t, err := parser.ValidateMultipartUploadType(v.Type, v.RawTag)

		if err != nil || t == parser.UploadNone {
			continue
		}

		if jsonTag == nil {
			jsonTag = &structtag.Tag{
				Key:  "json",
				Name: "-",
			}

			_ = tags.Set(jsonTag)
		} else {
			jsonTag.Name = "-"
		}

		if t == parser.UploadSingleFile {
			v.Type = &types.Nullable{
				Inner: &types.Object{
					Name: fileHeaderName,
				},
			}
		} else {
			v.Type = &types.Nullable{
				Inner: &types.Array{
					Inner: &types.Object{
						Name: fileHeaderName,
					},
				},
			}
		}

		v.RawTag = tags.String()

		obj.Entries[k] = v
	}
}
