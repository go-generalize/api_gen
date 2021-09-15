// Package parser parses API definitions in Go
package parser

import (
	"fmt"
	"go/ast"
	goparser "go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/go-generalize/api_gen/pkg/agerrors"
	"github.com/go-generalize/api_gen/pkg/common"
	"github.com/go-generalize/api_gen/pkg/util"
	go2tsparser "github.com/go-generalize/go2ts/pkg/parser"
	go2tstypes "github.com/go-generalize/go2ts/pkg/types"
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

	// Make the root path empty
	gr.Path = ""
	gr.RawPath = ""
	gr.Placeholder = ""

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

func (p *parser) parseFile(dir, fileName string, fset *token.FileSet, file *ast.File) (map[string]*Endpoint, error) {
	eps := make(map[string]*Endpoint)

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

			ep := eps[me]
			if ep == nil {
				ep = &Endpoint{}
				eps[me] = ep
			}

			rawPath := me[len(method):]
			ep.Method = method
			ep.RawPath = rawPath
			ep.Path = strcase.ToSnake(rawPath)
			ep.File = fileName

			if strings.HasPrefix(filepath.Base(fileName), "0_") {
				ep.Placeholder = stem(fileName)[2:]
			}

			if isRequest {
				ep.RequestPayload = structType
				ep.RequestPayloadName = name
				ep.RequestPos = typeSpec.Pos()
				ep.RequestLine = fset.Position(typeSpec.Pos()).Line
			} else {
				ep.ResponsePayload = structType
				ep.ResponsePayloadName = name
				ep.ResponsePos = typeSpec.Pos()
				ep.ResponseLine = fset.Position(typeSpec.Pos()).Line
			}
		}
	}

	type posEndpoint struct {
		pos token.Pos
		ep  *Endpoint
	}

	peps := make([]*posEndpoint, 0, len(eps)*2)
	for k, v := range eps {
		if v.RequestPayload == nil || v.ResponsePayload == nil {
			return nil, xerrors.Errorf("missing declaration for %s", k)
		}

		peps = append(peps, &posEndpoint{
			pos: v.RequestPos,
			ep:  v,
		})
		peps = append(peps, &posEndpoint{
			pos: v.ResponsePos,
			ep:  v,
		})
	}
	sort.Slice(peps, func(i, j int) bool {
		return peps[i].pos < peps[j].pos
	})

	for _, gr := range file.Comments {
		for _, cmnt := range gr.List {
			txt := cmnt.Text

			txt = strings.TrimPrefix(txt, "//")
			txt = strings.TrimSpace(txt)

			if !strings.HasPrefix(txt, "@") {
				continue
			}

			pos := cmnt.Pos()
			var ep *Endpoint
			for i := range peps {
				if peps[i].pos > pos {
					ep = peps[i].ep
					break
				}
			}

			if ep != nil {
				ep.SwagComments = append(ep.SwagComments, "// "+txt)
			}
		}
	}

	return eps, nil
}

func (p *parser) parseGo2tsDir(dir string) (map[string]go2tstypes.Type, error) {
	psr, err := go2tsparser.NewParser(dir, func(fo *go2tsparser.FilterOpt) bool {
		if !fo.BasePackage {
			return false
		}

		return common.ParserFilter(fo)
	})

	if err != nil {
		return nil, xerrors.Errorf("failed to parse %s: %w", dir, err)
	}

	parsed, err := psr.Parse()

	if err != nil {
		return nil, xerrors.Errorf("failed to parse %s: %w", dir, err)
	}

	results := make(map[string]go2tstypes.Type)
	for key, pkg := range parsed {
		if !(strings.HasSuffix(key, "Request") || strings.HasSuffix(key, "Response")) {
			continue
		}

		idx := strings.LastIndex(key, ".")
		if idx == -1 {
			continue
		}

		name := key[idx+1:]

		results[name] = pkg
	}

	return results, nil
}

func (p *parser) parsePackage(dir string) (*Group, error) {
	gomod, err := p.getGoModulePackage(dir)

	if err != nil {
		return nil, xerrors.Errorf("failed to get go module package: %w", err)
	}

	fset := token.NewFileSet()

	endpoints := map[string]*Endpoint{}
	gr := &Group{
		Dir:        dir,
		ImportPath: gomod,
		RawPath:    filepath.Base(dir),
		Path:       filepath.Base(dir),
	}

	if strings.HasPrefix(gr.RawPath, "_") {
		gr.Placeholder = gr.RawPath[1:]
	}

	fifos, err := os.ReadDir(dir)

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
		child.parentGroup = gr

		gr.Children = append(gr.Children, child)
	}

	sort.Slice(gr.Children, func(i, j int) bool {
		return gr.Children[i].RawPath < gr.Children[j].RawPath
	})

	hasGoFile := false
	for _, fifo := range fifos {
		if filepath.Ext(fifo.Name()) == ".go" {
			hasGoFile = true
			break
		}
	}

	if !hasGoFile {
		return gr, nil
	}

	go2tsTypes, err := p.parseGo2tsDir(dir)

	if err != nil {
		return nil, xerrors.Errorf("failed to parse dir with go2ts parser: %w", err)
	}

	pkgs, err := goparser.ParseDir(fset, dir, nil, goparser.AllErrors|goparser.ParseComments)

	if err != nil {
		return nil, xerrors.Errorf("failed to parse dir(%s): %w", dir, err)
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

			eps, err := p.parseFile(dir, name, fset, file)

			if err != nil {
				return nil, xerrors.Errorf("failed to parse %s in %s: %w", name, dir, err)
			}

			for k, v := range eps {
				endpoints[k] = v
			}
		}

		gr.Endpoints = make([]*Endpoint, 0, len(endpoints))

		for _, v := range endpoints {
			v := v
			if v.RequestPayload == nil ||
				v.ResponsePayload == nil {
				continue
			}
			v.parentGroup = gr

			var err error
			v.GetFullPath("", func(rawPath, path, placeholder string) string {
				if placeholder != "" {
					queryParamField := util.FindStructField(v.RequestGo2tsPayload, QueryParamTag, placeholder)
					if queryParamField == "" {
						err = agerrors.NewParserError(
							v.File,
							v.RequestLine,
							fmt.Sprintf(
								"'%s' is missing in %s as field name or param tag",
								placeholder, v.RequestPayloadName,
							),
						)
					}
				}

				return ""
			})

			if err != nil {
				return nil, xerrors.Errorf("invalid endpoint: %w", err)
			}

			req, ok := go2tsTypes[v.RequestPayloadName]
			if !ok {
				return nil, xerrors.Errorf("request type is not found from types parsed by go2ts: %s", v.RequestPayloadName)
			}
			reqObj, ok := req.(*go2tstypes.Object)
			if !ok {
				return nil, xerrors.Errorf("request type is not object: %s", v.RequestPayloadName)
			}
			v.RequestGo2tsPayload = reqObj

			res, ok := go2tsTypes[v.ResponsePayloadName]
			if !ok {
				return nil, xerrors.Errorf("response type is not found from types parsed by go2ts: %s", v.ResponsePayloadName)
			}
			resObj, ok := res.(*go2tstypes.Object)
			if !ok {
				return nil, xerrors.Errorf("response type is not object: %s", v.ResponsePayloadName)
			}
			v.ResponseGo2tsPayload = resObj

			gr.Endpoints = append(gr.Endpoints, v)
		}

		sort.Slice(gr.Endpoints, func(i, j int) bool {
			return string(gr.Endpoints[i].Method)+gr.Endpoints[i].RawPath < string(gr.Endpoints[j].Method)+gr.Endpoints[j].RawPath
		})
	}

	return gr, nil
}
