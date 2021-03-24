package server

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/iancoleman/strcase"
	"golang.org/x/xerrors"
)

func (g *Generator) generateController(root, propsPackage string, ep *parser.Endpoint) error {
	rel := ep.GetParent().GetFullPath("/", func(rawPath, path, placeholder string) string {
		return rawPath
	})
	if err := os.MkdirAll(filepath.Join(root, rel), 0777); err != nil {
		return xerrors.Errorf("failed to mkdir %s: %w", filepath.Join(root, rel), err)
	}

	filename := fmt.Sprintf("%s_%s.go", strings.ToLower(string(ep.Method)), strcase.ToSnake(ep.Path))
	if len(ep.Path) == 0 {
		filename = fmt.Sprintf("%s.go", strings.ToLower(string(ep.Method)))
	}

	path := filepath.Join(root, rel, filename)

	if _, err := os.Stat(path); err == nil {
		return nil
	}

	fp, err := os.Create(path)
	if err != nil {
		return xerrors.Errorf("failed to create %s: %w", path, err)
	}
	defer fp.Close()

	handler := strcase.ToCamel(strings.ToLower(string(ep.Method)) + ep.RawPath)
	if err := g.controllerTemplate.Execute(fp, struct {
		*parser.Endpoint
		ControllerName                        string
		ControllerInterfaceName               string
		HandlerName                           string
		AppVersion                            string
		ControllerPropsPackage                string
		RequestStructName, ResponseStructName string
		TypesPackage                          string
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
	}); err != nil {
		return xerrors.Errorf("failed to generate controller: %w", err)
	}

	return nil
}
