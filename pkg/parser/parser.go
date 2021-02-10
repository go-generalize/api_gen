package parser

import (
	"go/ast"
	goparser "go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/go-utils/gopackages"
	"github.com/iancoleman/strcase"
	"golang.org/x/xerrors"
)

// Parse parses
func Parse(dir string) (*Group, error) {
	parser, err := newParser(dir)

	if err != nil {
		return nil, xerrors.Errorf("failed to initialize parser: %w", err)
	}

	gr, err := parser.parsePackage(dir)

	if err != nil {
		return nil, xerrors.Errorf("failed to parse package: %w", err)
	}

	return gr, nil
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
	abs, err := filepath.Abs(dir)

	if err != nil {
		return "", xerrors.Errorf("failed to get absolute path for %s: %w", dir, abs)
	}

	d, err := filepath.Rel(p.moduleDir, abs)

	if err != nil {
		return "", xerrors.Errorf("failed to get relative path from module root: %w", err)
	}

	return filepath.Clean(d), nil
}

func (p *parser) getGoModulePackage(dir string) (string, error) {
	rel, err := p.relativePathFromRoot(dir)

	if err != nil {
		return "", xerrors.Errorf("failed to get relative path for go module: %w", err)
	}

	return filepath.Join(p.module, rel), nil
}

func (p *parser) parseFile(dir, fileName string, fset *token.FileSet, file *ast.File, endpoints map[string]*Endpoint) {
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			// 型定義
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			name := typeSpec.Name.Name

			method := getMethodType(name)
			if method == "" {
				continue
			}

			var me string
			var isRequest bool
			if strings.HasSuffix(name, "Request") {
				me = strings.TrimSuffix(name, "Request")
				isRequest = true
			} else if strings.HasSuffix(name, "Response") {
				me = strings.TrimSuffix(name, "Response")
				isRequest = false
			} else {
				continue
			}

			ep, found := endpoints[me]

			if !found {
				ep = &Endpoint{}
				endpoints[me] = ep
			}

			rawPath := me[len(method):]
			ep.Method = method
			ep.RawPath = rawPath
			ep.Path = strcase.ToSnake(rawPath)

			if strings.HasPrefix(filepath.Base(fileName), "0_") {
				ep.Placeholder = strings.ReplaceAll(stem(fileName), "_id", "ID")[2:]
			}

			if isRequest {
				ep.RequestPayload = structType
			} else {
				ep.ResponsePayload = structType
			}
		}
	}
}

func (p *parser) parsePackage(dir string) (*Group, error) {
	gomod, err := p.getGoModulePackage(dir)

	if err != nil {
		return nil, xerrors.Errorf("failed to get go module package: %w", err)
	}

	fset := token.NewFileSet()

	pkgs, err := goparser.ParseDir(fset, dir, nil, goparser.AllErrors)

	if err != nil {
		return nil, xerrors.Errorf("failed to parse dir(%s): %w", dir, err)
	}

	endpoints := map[string]*Endpoint{}
	gr := &Group{
		Dir:        dir,
		ImportPath: gomod,
		RawPath:    filepath.Base(dir),
		Path:       strcase.ToSnake(filepath.Base(dir)),
	}

	if strings.HasPrefix(gr.RawPath, "_") {
		gr.Placeholder = gr.RawPath[1:]
	}

	for name, v := range pkgs {
		if strings.HasSuffix(name, "_test") {
			continue
		}

		for name, file := range v.Files {
			stem := strings.TrimSuffix(filepath.Base(name), filepath.Ext(name))

			if strings.HasSuffix(stem, "_test") {
				continue
			}

			p.parseFile(dir, name, fset, file, endpoints)
		}

		gr.Endpoints = make([]*Endpoint, 0, len(endpoints))

		for _, v := range endpoints {
			if v.RequestPayload == nil ||
				v.ResponsePayload == nil {
				continue
			}

			gr.Endpoints = append(gr.Endpoints, v)
		}

		continue
	}

	fifos, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, xerrors.Errorf("failed to list all child files/directories: %w", err)
	}

	for _, fifo := range fifos {
		if !fifo.IsDir() {
			continue
		}

		child, err := p.parsePackage(filepath.Join(dir, fifo.Name()))

		if err != nil {
			return nil, xerrors.Errorf("failed to parse package for %s: %w", fifo.Name(), err)
		}

		if len(child.Endpoints) == 0 && len(child.Children) == 0 {
			continue
		}

		gr.Children = append(gr.Children, child)
	}

	return gr, nil
}
