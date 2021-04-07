// Package server generates server-side library
package server

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"go/format"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"

	"golang.org/x/xerrors"
)

type bundlerEndpoint struct {
	ControllerImport      string
	ControllerInitializer string
	TypesImport           string
	Method                string
	Path                  string
	RequestStructName     string
	HandlerName           string
}

type bundlerEndpointForImports struct {
	bundlerEndpoint
	ControllerImportAlias string
	TypesImportAlias      string
}

type bundlerImport struct {
	Alias, Path string
}

type bundlerData struct {
	Endpoints  []bundlerEndpointForImports
	Imports    []bundlerImport
	AppVersion string
}

func (g *Generator) generateControllerBundler(
	endpoints []*bundlerEndpoint,
	propsPackage string,
	apierrorPath string,
) error {
	path := filepath.Join(g.base, "bundler.go")

	fp, err := os.Create(path)

	if err != nil {
		return xerrors.Errorf("failed to create %s: %w", path, err)
	}
	defer fp.Close()

	bd := &bundlerData{
		AppVersion: g.AppVersion,
	}

	bd.Endpoints = make([]bundlerEndpointForImports, 0, len(endpoints))
	imports := make(map[string]string)
	for i := range endpoints {
		ctrlAlias := controllerImportAlias(endpoints[i].ControllerImport)
		typesAlias := typesImportAlias(endpoints[i].TypesImport)

		bd.Endpoints = append(bd.Endpoints, bundlerEndpointForImports{
			bundlerEndpoint:       *endpoints[i],
			ControllerImportAlias: ctrlAlias,
			TypesImportAlias:      typesAlias,
		})

		imports[endpoints[i].ControllerImport] = ctrlAlias
		imports[endpoints[i].TypesImport] = typesAlias
	}

	imports[propsPackage] = "props"
	imports["golang.org/x/xerrors"] = "xerrors"
	imports["github.com/labstack/echo/v4"] = "echo"
	imports[apierrorPath] = "apierror"

	bd.Imports = make([]bundlerImport, 0, len(imports))
	for k, v := range imports {
		bd.Imports = append(bd.Imports, bundlerImport{
			Alias: v,
			Path:  k,
		})
	}
	sort.Slice(bd.Imports, func(i, j int) bool {
		return bd.Imports[i].Path < bd.Imports[j].Path
	})

	buf := bytes.NewBuffer(nil)
	err = g.controllerBundlerTemplate.Execute(buf, bd)

	if err != nil {
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

func typesImportAlias(p string) string {
	return "types_" + path.Base(p) + "_" + hash(p)
}

func controllerImportAlias(p string) string {
	return "ctrl_" + path.Base(p) + "_" + hash(p)
}

func hash(path string) string {
	hash := sha1.Sum([]byte(path))

	return hex.EncodeToString(hash[:])[:8]
}
