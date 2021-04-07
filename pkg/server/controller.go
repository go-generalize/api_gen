package server

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/iancoleman/strcase"
	"golang.org/x/xerrors"
)

func (g *Generator) generateController(root, propsPackage string, ep *parser.Endpoint) (*bundlerEndpoint, error) {
	rel := ep.GetParent().GetFullPath("/", func(rawPath, path, placeholder string) string {
		return rawPath
	})
	dir := filepath.Join(root, rel)

	if err := os.MkdirAll(dir, 0777); err != nil {
		return nil, xerrors.Errorf("failed to mkdir %s: %w", filepath.Join(root, rel), err)
	}

	filename := fmt.Sprintf("%s_%s.go", strings.ToLower(string(ep.Method)), strcase.ToSnake(ep.Path))
	if len(ep.Path) == 0 {
		filename = fmt.Sprintf("%s.go", strings.ToLower(string(ep.Method)))
	}

	controllerImportPath, err := g.module.GetImportPath(dir)
	if err != nil {
		return nil, xerrors.Errorf("failed to get import path for dir: %w", err)
	}

	handler := strcase.ToCamel(strings.ToLower(string(ep.Method)) + ep.RawPath)
	be := &bundlerEndpoint{
		ControllerImport:      controllerImportPath,
		ControllerInitializer: "New" + handler + "Controller",
		TypesImport:           ep.GetParent().ImportPath,
		Method:                string(ep.Method),
		Path: ep.GetFullPath("/", func(rawPath, path, placeholder string) string {
			if placeholder != "" {
				return ":" + placeholder
			}

			return path
		}),
		RequestStructName: ep.RequestPayloadName,
		HandlerName:       handler,
	}

	path := filepath.Join(root, rel, filename)

	if _, err := os.Stat(path); err == nil {
		return be, nil
	}

	fp, err := os.Create(path)
	if err != nil {
		return nil, xerrors.Errorf("failed to create %s: %w", path, err)
	}
	defer fp.Close()

	buf := bytes.NewBuffer(nil)
	if err := g.controllerTemplate.Execute(buf, struct {
		*parser.Endpoint
		ControllerName                        string
		ControllerInterfaceName               string
		HandlerName                           string
		AppVersion                            string
		ControllerPropsPackage                string
		RequestStructName, ResponseStructName string
		TypesPackage                          string
		FullPath                              string
	}{
		Endpoint:                ep,
		ControllerName:          strcase.ToLowerCamel(handler) + "Controller",
		ControllerInterfaceName: handler + "Controller",
		HandlerName:             handler,
		AppVersion:              g.AppVersion,
		RequestStructName:       ep.RequestPayloadName,
		ResponseStructName:      ep.ResponsePayloadName,
		ControllerPropsPackage:  propsPackage,
		TypesPackage:            ep.GetParent().ImportPath,
		FullPath:                rel,
	}); err != nil {
		return nil, xerrors.Errorf("failed to generate controller: %w", err)
	}

	src, err := format.Source(buf.Bytes())

	if err != nil {
		return nil, xerrors.Errorf("failed to format code: %w", err)
	}

	if _, err := io.Copy(fp, bytes.NewBuffer(src)); err != nil {
		return nil, xerrors.Errorf("failed to write %s: %w", path, err)
	}

	return be, nil
}
