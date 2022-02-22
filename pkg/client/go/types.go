// Package clientgo generates client-side Go library
package clientgo

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/structtag"
	"github.com/go-generalize/api_gen/v2/pkg/parser"
	"github.com/go-generalize/go-easyparser/types"
	"github.com/go-generalize/go2go"
	"golang.org/x/xerrors"
)

const headerComment = `// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: %s

`

// GenerateTypes generates request/response types in Go
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

	replaceFileHeaderAll(gr.ParsedTypes)

	endpointStructs := make([]string, 0, 20)
	for _, ep := range gr.Endpoints {
		endpointStructs = append(endpointStructs, gr.ImportPath+"."+ep.RequestPayloadName)
		endpointStructs = append(endpointStructs, gr.ImportPath+"."+ep.ResponsePayloadName)
	}

	gen := go2go.NewGenerator(gr.ParsedTypes, endpointStructs)

	gen.ExternalGenerator = func(t types.Type) (*go2go.GeneratedType, bool) {
		o, ok := t.(*types.Object)

		if !ok {
			return nil, false
		}

		if o.Name == fileHeaderName {
			return &go2go.GeneratedType{
				Path: "github.com/go-generalize/multipart-util",
				Name: "MultipartPayload",
			}, true
		}

		return nil, false
	}

	code, err := gen.Generate()

	if err != nil {
		return xerrors.Errorf("failed to initialize new go2go generator: %w", err)
	}

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

const fileHeaderName = "mime/multipart.FileHeader"

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

		tags, err := structtag.Parse(v.RawTag)

		if err != nil {
			continue
		}

		jsonTag, err := tags.Get("json")

		if err != nil {
			jsonTag = &structtag.Tag{
				Key:  "json",
				Name: "-",
			}
		}
		tags.Set(jsonTag)

		t, err := parser.ValidateMultipartUploadType(v.Type, v.RawTag)

		if err != nil || t == parser.UploadNone {
			continue
		}

		if t == parser.UploadSingleFile {
			v.Type = &types.Nullable{
				Inner: &types.Object{
					Name: fileHeaderName,
				},
			}
		} else {
			v.Type = &types.Array{
				Inner: &types.Nullable{
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
