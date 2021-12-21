// Package parser parses API definitions in Go
package parser

import (
	"fmt"
	"go/ast"
	goparser "go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/go-generalize/api_gen/v2/pkg/agerrors"
	"github.com/go-generalize/api_gen/v2/pkg/common"
	"github.com/go-generalize/api_gen/v2/pkg/util"
	go2tsparser "github.com/go-generalize/go2ts/pkg/parser"
	go2tstypes "github.com/go-generalize/go2ts/pkg/types"
	"github.com/go-utils/gopackages"
	"github.com/iancoleman/strcase"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
)

// Parse parses
func Parse(dir string) (*Group, error) {
	parser, err := newParser(dir)

	if err != nil {
		return nil, xerrors.Errorf("failed to initialize parser: %w", err)
	}

	dir, err = filepath.Abs(dir)

	if err != nil {
		return nil, xerrors.Errorf("failed to get absolute path for %s: %w", dir, err)
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

	threadLock chan struct{}
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

	psr := &parser{
		module:    module,
		moduleDir: moduleDir,
		rootDir:   dir,
	}
	psr.SetParallelism(runtime.NumCPU())

	return psr, nil
}

// SetParallelism limits the max number of threads
// Default: runtime.NumCPU()
func (p *parser) SetParallelism(num int) {
	threadLock := make(chan struct{}, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		threadLock <- struct{}{}
	}

	p.threadLock = threadLock
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

func (p *parser) extractReqRespTypes(parsed map[string]go2tstypes.Type, basePackage string) map[string]go2tstypes.Type {
	results := make(map[string]go2tstypes.Type)
	for key, pkg := range parsed {
		if !strings.HasPrefix(key, basePackage+".") {
			continue
		}

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

	return results
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

	eg := errgroup.Group{}

	for _, fifo := range fifos {
		if !fifo.IsDir() {
			continue
		}
		fifo := fifo
		idx := len(gr.Children)
		gr.Children = append(gr.Children, nil)

		eg.Go(func() error {
			child, err := p.parsePackage(filepath.Join(dir, fifo.Name()))

			if err != nil {
				return xerrors.Errorf("failed to parse package for %s: %w", fifo.Name(), err)
			}

			child.parentGroup = gr
			gr.Children[idx] = child
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, xerrors.Errorf("failed to parse children of %s: %w", dir, err)
	}

	<-p.threadLock
	go func() {
		p.threadLock <- struct{}{}
	}()

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

	psr, err := go2tsparser.NewParser(dir, common.ParserFilter)

	if err != nil {
		return nil, xerrors.Errorf("failed to initialize %s: %w", dir, err)
	}

	parsed, err := psr.Parse()

	if err != nil {
		return nil, xerrors.Errorf("failed to parse %s: %w", dir, err)
	}

	gr.ParsedTypes = parsed

	go2tsTypes := p.extractReqRespTypes(parsed, psr.GetBasePackage())

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
				v.parentGroup = gr
				endpoints[k] = v
			}
		}

		gr.Endpoints = make([]*Endpoint, 0, len(endpoints))

		for _, v := range endpoints {
			err = p.updateEndpoint(v, go2tsTypes)

			if err == errNilReqRespPayload {
				continue
			}

			if err != nil {
				return nil, xerrors.Errorf("failed to verity endpoint %s: %w", v.RawPath, err)
			}

			gr.Endpoints = append(gr.Endpoints, v)
		}

		sort.Slice(gr.Endpoints, func(i, j int) bool {
			return string(gr.Endpoints[i].Method)+gr.Endpoints[i].RawPath < string(gr.Endpoints[j].Method)+gr.Endpoints[j].RawPath
		})
	}

	return gr, nil
}

var (
	errNilReqRespPayload = fmt.Errorf("request or response payload is nil")
)

func (p *parser) updateEndpoint(v *Endpoint, reqRespTypes map[string]go2tstypes.Type) error {
	if v.RequestPayload == nil ||
		v.ResponsePayload == nil {
		return nil
	}

	req, ok := reqRespTypes[v.RequestPayloadName]
	if !ok {
		return xerrors.Errorf("request type is not found from types parsed by go2ts: %s", v.RequestPayloadName)
	}
	reqObj, ok := req.(*go2tstypes.Object)
	if !ok {
		return xerrors.Errorf("request type is not object: %s", v.RequestPayloadName)
	}
	v.RequestGo2tsPayload = reqObj

	res, ok := reqRespTypes[v.ResponsePayloadName]
	if !ok {
		return xerrors.Errorf("response type is not found from types parsed by go2ts: %s", v.ResponsePayloadName)
	}
	resObj, ok := res.(*go2tstypes.Object)
	if !ok {
		return xerrors.Errorf("response type is not object: %s", v.ResponsePayloadName)
	}
	v.ResponseGo2tsPayload = resObj

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
		return xerrors.Errorf("invalid endpoint: %w", err)
	}

	return nil
}
