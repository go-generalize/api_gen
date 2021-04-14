package server

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"path/filepath"
	"strings"

	go2json "github.com/go-generalize/api_gen/pkg/go2json"
	"github.com/go-generalize/api_gen/pkg/parser"
	go2tsp "github.com/go-generalize/go2ts/pkg/parser"
	go2tst "github.com/go-generalize/go2ts/pkg/types"
	"github.com/iancoleman/strcase"
	"golang.org/x/xerrors"
)

func (g *Generator) generateMock(apierrorPackage, propsPackage string) error {
	mockPath := filepath.Join(g.base, "mock")
	mockJsonPath := filepath.Join(mockPath, "json")
	mockControllerPath := filepath.Join(mockPath, "controller")

	mockPackage, err := g.module.GetImportPath(mockPath)
	if err != nil {
		return xerrors.Errorf("failed to get import path for %s: %w", mockPath, err)
	}

	types, err := g.parsePackage()

	if err != nil {
		return xerrors.Errorf("failed to parse package: %w", err)
	}

	if err := os.MkdirAll(mockJsonPath, 0777); err != nil {
		return xerrors.Errorf("failed to create %s: %w", mockPath, err)
	}

	// mock/mock.go
	if err := g.generateMockJsonFS(mockPath); err != nil {
		return xerrors.Errorf("failed to generate mock.go: %w", err)
	}

	// mock/json/**/*.json
	if err := go2json.NewGenerator(types).Generate(mockJsonPath); err != nil {
		return xerrors.Errorf("failed to generate mock json: %w", err)
	}

	// mock/controller/**/*.go
	be, err := g.generateControllers(
		mockControllerPath,
		g.group,
		func(ep *parser.Endpoint) (*bundlerEndpoint, error) {
			return g.generateMockController(mockControllerPath, mockPackage, ep)
		},
	)

	if err != nil {
		return xerrors.Errorf("failed to generate mock/controller: %w", err)
	}

	// mock_bundler.go
	if err := g.generateControllerBundler(
		be,
		propsPackage,
		apierrorPackage,
		"mock_bundler.go",
		true,
	); err != nil {
		return xerrors.Errorf("failed to generate mock_bundler.go: %w", err)
	}

	return nil
}

func (g *Generator) parsePackage() (map[string]go2tst.Type, error) {
	psr, err := go2tsp.NewParser(g.group.Dir, func(opt *go2tsp.FilterOpt) bool {
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
	})

	if err != nil {
		return nil, xerrors.Errorf("failed to initializer go2ts parser: %w", err)
	}

	types, err := psr.Parse()

	if err != nil {
		return nil, xerrors.Errorf("failed to parse %s: %w", g.group.Dir)
	}

	return types, nil
}

func (g *Generator) generateMockController(root, mockPackage string, ep *parser.Endpoint) (*bundlerEndpoint, error) {
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
	if err := g.mockControllerTemplate.Execute(buf, struct {
		*parser.Endpoint
		ControllerName                        string
		ControllerInterfaceName               string
		HandlerName                           string
		AppVersion                            string
		ControllerPropsPackage                string
		RequestStructName, ResponseStructName string
		TypesPackage                          string
		FullPath                              string
		MockPackage                           string
		RawEndpointPath                       string
		JsonDir                               string
	}{
		Endpoint:                ep,
		ControllerName:          strcase.ToLowerCamel(handler) + "Controller",
		ControllerInterfaceName: handler + "Controller",
		HandlerName:             handler,
		AppVersion:              g.AppVersion,
		RequestStructName:       ep.RequestPayloadName,
		ResponseStructName:      ep.ResponsePayloadName,
		TypesPackage:            ep.GetParent().ImportPath,
		FullPath:                rel,
		MockPackage:             mockPackage, // TODO: Update mock package
		RawEndpointPath: ep.GetFullPath("/", func(rawPath, path, placeholder string) string {
			return path
		}),
		JsonDir: ep.GetParent().GetFullPath("/", func(rawPath, path, placeholder string) string {
			if placeholder != "" {
				return placeholder
			}
			return path
		}) + "/" + strings.ToLower(string(ep.Method)) + "_" + ep.Path,
	}); err != nil {
		return nil, xerrors.Errorf("failed to generate controller: %w", err)
	}

	src, err := format.Source(buf.Bytes())

	if err != nil {
		return nil, xerrors.Errorf("failed to format code(%s): %w", buf.String(), err)
	}

	if _, err := io.Copy(fp, bytes.NewBuffer(src)); err != nil {
		return nil, xerrors.Errorf("failed to write %s: %w", path, err)
	}

	return be, nil
}

func (g *Generator) generateMockJsonFS(mockPath string) error {
	buf := bytes.NewBuffer(nil)

	file := filepath.Join(mockPath, "mock.go")
	fp, err := os.Create(file)

	if err != nil {
		return xerrors.Errorf("failed to create %s: %w", file, err)
	}
	defer fp.Close()

	if err := g.mockTemplate.Execute(buf, map[string]string{
		"AppVersion": g.AppVersion,
	}); err != nil {
		return xerrors.Errorf("failed to execute template: %w", err)
	}

	src, err := format.Source(buf.Bytes())

	if err != nil {
		return xerrors.Errorf("failed to format code: %w", err)
	}

	if _, err := io.Copy(fp, bytes.NewBuffer(src)); err != nil {
		return xerrors.Errorf("failed to write apierror.go: %w", err)
	}

	return nil
}
