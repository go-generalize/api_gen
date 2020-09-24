package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-generalize/api_gen/common"

	_ "github.com/go-generalize/api_gen/client_generator/statik"
	go2tsgenerator "github.com/go-generalize/go2ts/pkg/generator"
	go2tsparser "github.com/go-generalize/go2ts/pkg/parser"
	"github.com/iancoleman/strcase"
)

var supportedMethods = []string{
	"get",
	"post",
	"put",
	"delete",
	"patch",
}

type endpoint struct {
	methodEndpoint string
	path           string
	rawName        string
	fileName       string

	requestStructObject  *ast.StructType
	responseStructObject *ast.StructType

	method            string
	request, response bool
}

type pkgParser struct {
	endpoints map[string]*endpoint

	structs []string
}

func newPkgParser() *pkgParser {
	return &pkgParser{
		endpoints: map[string]*endpoint{},
	}
}

func (p *pkgParser) parseFile(pathName, dir string, fset *token.FileSet, file *ast.File) {
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok != token.TYPE {
			continue
		}
		packageName := file.Name.Name

		for _, spec := range genDecl.Specs {
			// 型定義
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			name := typeSpec.Name.Name

			lowered := strings.ToLower(name)
			method := ""
			for i := range supportedMethods {
				if strings.HasPrefix(lowered, supportedMethods[i]) {
					method = supportedMethods[i]

					break
				}
			}

			if len(method) == 0 {
				continue
			}

			var me string
			if strings.HasSuffix(name, "Request") {
				me = strings.TrimSuffix(name, "Request")
			} else if strings.HasSuffix(name, "Response") {
				me = strings.TrimSuffix(name, "Response")
			} else {
				continue
			}

			goFilePath := fset.File(spec.Pos()).Name()
			goFileName := filepath.Base(goFilePath)
			goFileName = goFileName[:len(goFileName)-len(filepath.Ext(goFileName))]

			if _, ok := p.endpoints[me]; !ok {
				p.endpoints[me] = &endpoint{}
			}

			p.endpoints[me].fileName = goFilePath
			p.endpoints[me].method = method
			p.endpoints[me].methodEndpoint = me

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				fmt.Printf("\x1b[31mskip: %s \n"+
					"  must be Struct.\x1b[0m\n", fset.Position(typeSpec.Type.Pos()).String())
				continue
			}

			if strings.HasSuffix(name, "Request") {
				p.endpoints[me].request = true

				p.endpoints[me].rawName = strings.TrimSuffix(name, "Request")
				p.endpoints[me].path = path.Join(pathName, strcase.ToSnake(strings.TrimSuffix(name[len(method):], "Request")))
				p.endpoints[me].requestStructObject = structType
			} else {
				p.endpoints[me].response = true

				p.endpoints[me].rawName = strings.TrimSuffix(name, "Response")
				p.endpoints[me].path = path.Join(pathName, strcase.ToSnake(strings.TrimSuffix(name[len(method):], "Response")))
				p.endpoints[me].responseStructObject = structType
			}

			if strings.HasPrefix(goFileName, "0_") {
				goFileName = goFileName[2:]
				endpointPath := p.endpoints[me].path
				if !strings.HasPrefix(strings.ToLower(goFileName), ":id") {
					goFileName = strings.ReplaceAll(goFileName, "_id", "ID")
				}
				endpointPath = strcase.ToCamel(goFileName)
				endpointPath = fmt.Sprintf("/_%s", strings.ToLower(string(endpointPath[0]))+endpointPath[1:])
				p.endpoints[me].path = filepath.Join(filepath.Dir(p.endpoints[me].path), endpointPath)
			}

			packageNameFromFilePath := filepath.Base(dir)
			if packageName != packageNameFromFilePath {
				fmt.Printf("\x1b[31mskip: %s/.%s \n"+
					"  The file path and the actual package name must match.\n"+
					"  package name=%s, require=%s\x1b[0m\n",
					p.endpoints[me].path, p.endpoints[me].rawName, packageNameFromFilePath, packageName)
				continue
			}

			p.structs = append(p.structs, name)
		}
	}
}

func (p *pkgParser) parsePackage(pathName, dir string, fset *token.FileSet, pkg *ast.Package) {
	for name, file := range pkg.Files {
		stem := strings.TrimSuffix(filepath.Base(name), filepath.Ext(name))

		if strings.HasSuffix(stem, "_test") {
			continue
		}

		p.parseFile(pathName, dir, fset, file)
	}
}

func (p *pkgParser) parseDir(pathName, dir string) {
	var fset = token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, dir, nil, parser.AllErrors)

	if err != nil {
		log.Fatalf("failed to parse dir: %+v", err)
	}

	for name, v := range pkgs {
		if strings.HasSuffix(name, "_test") {
			continue
		}

		p.parsePackage(pathName, dir, fset, v)
	}
}

func walk(p, url string, generator *clientGenerator, parent *clientType) {
	endpointReplaceMatchRule := regexp.MustCompile(`(?m):(.*?)(/|$)`)
	pkgParser := newPkgParser()

	pkgParser.parseDir(url, p)

	for k, v := range pkgParser.endpoints {
		fmt.Println(k, *v)
	}

	parent.Name = strcase.ToCamel(strings.ReplaceAll(url, "/", "-")) + "Client"

	pairs := make([]importPair, 0, len(pkgParser.structs))
	for i := range pkgParser.structs {
		pairs = append(pairs, importPair{
			Name:   pkgParser.structs[i],
			NameAs: strcase.ToCamel(strings.ReplaceAll(url+"/"+pkgParser.structs[i], "/", "-")),
		})
	}
	if len(pairs) != 0 {
		generator.Imports = append(
			generator.Imports,
			importType{
				Path:  "./classes" + url + "/types",
				Pairs: pairs,
			},
		)
	}

	for _, ep := range pkgParser.endpoints {
		if !(ep.request && ep.response) {
			continue
		}

		replaced := strcase.ToCamel(strings.ReplaceAll(url+"/"+ep.rawName, "/", "-"))
		urlParams := make([]string, 0)
		endpointPath := ep.path
		endpointPath = strings.ReplaceAll(endpointPath, "/_", "/:")
		requestStruct := ep.requestStructObject
		endpointPath = endpointReplaceMatchRule.ReplaceAllStringFunc(endpointPath, func(s string) string {
			hasSuffixSlash := strings.HasSuffix(s, "/")
			param := ""
			if hasSuffixSlash {
				param = s[1 : len(s)-1]
			} else {
				param = s[1:]
			}

			st := requestStruct
			fieldList := st.Fields.List
			for _, fields := range fieldList {
				if fields.Tag == nil {
					continue
				}
				tags := reflect.StructTag(strings.Trim(fields.Tag.Value, "`"))

				jsonTag, jsonOk := tags.Lookup("json")

				if paramTag, ok := tags.Lookup("param"); ok {
					if paramTag == param {
						if jsonOk {
							param = jsonTag
						} else {
							param = fields.Names[0].Name
						}
					}
				}
			}

			urlParams = append(urlParams, param)

			suffix := ""
			if hasSuffixSlash {
				suffix = "/"
			}
			return fmt.Sprintf("${encodeURI(param.%s.toString())}%s", param, suffix)
		})

		responseFields := ep.responseStructObject.Fields

		parent.Methods = append(
			parent.Methods,
			&methodType{
				Name:         strcase.ToLowerCamel(ep.methodEndpoint),
				RequestType:  replaced + "Request",
				ResponseType: replaced + "Response",
				Method:       strings.ToUpper(ep.method),
				Endpoint:     endpointPath,
				URLParams:    urlParams,
				HasFields:    responseFields.List != nil && len(responseFields.List) > 0,
			},
		)
	}

	classesDir := filepath.Join(generator.OutputDir, "./classes/")

	if len(pkgParser.structs) != 0 {
		if err := os.MkdirAll(classesDir+url, 0774); err != nil {
			log.Fatalf("failed to MkdirAll: %+v", err)
		}
	}

	parser, err := go2tsparser.NewParser(p)

	if err != nil {
		log.Fatalf("failed to initialize go2ts parser: %+v", err)
	}
	parser.Filter = func(name string) bool {
		return strings.HasSuffix(name, "Request") || strings.HasSuffix(name, "Response")
	}

	types, err := parser.Parse()

	if err != nil {
		log.Fatalf("failed to parse go files: %+v", err)
	}

	ts := go2tsgenerator.NewGenerator(types).Generate()

	filename := filepath.Join(
		classesDir,
		fmt.Sprintf("/%s/types.ts", url),
	)

	if err := ioutil.WriteFile(filename, []byte(ts), os.ModePerm); err != nil {
		log.Fatalf("failed to write into %s: %+v", filename, err)
	}

	fifos, err := ioutil.ReadDir(p)

	if err != nil {
		log.Fatalf("failed to list files: %+v", err)
	}

	for i := range fifos {
		if !fifos[i].IsDir() {
			continue
		}

		client := new(clientType)

		nextURL := path.Join(url, fifos[i].Name())

		parent.Children = append(
			parent.Children,
			childrenType{
				Name:      fifos[i].Name(),
				ClassName: strcase.ToCamel(strings.ReplaceAll(nextURL, "/", "-")) + "Client",
			},
		)

		walk(filepath.Join(p, fifos[i].Name()), nextURL, generator, client)

		generator.ChildrenClients = append(generator.ChildrenClients, client)
	}
}

func main() {
	l := len(os.Args)
	if l < 2 {
		fmt.Println("You have to specify the path of target go file")
		os.Exit(1)
	}

	versionFlag := flag.Bool("v", false, "print version")
	outputDir := flag.String("o", "./", "output directory of generated codes")

	flag.Parse()

	if *versionFlag {
		fmt.Println(common.AppVersion)
		return
	}

	outputFullPath, err := filepath.Abs(*outputDir)
	if err != nil {
		log.Fatalf("failed to run filepath.Abs: %+v", err)
	}

	var stat os.FileInfo
	if stat, err = os.Stat(outputFullPath); err != nil {
		if err = os.MkdirAll(outputFullPath, 0774); err != nil {
			log.Fatalf("failed to MkdirAll: %+v", err)
		}
	} else if !stat.IsDir() {
		log.Fatalf("-o specified is not a directory")
	}

	log.Printf("output dir: %s", outputFullPath)
	classesDir := filepath.Join(outputFullPath, "./classes")

	if err = os.RemoveAll(classesDir); err != nil {
		log.Fatalf("failed to run RemoveAll: %+v", err)
	}
	if err = os.MkdirAll(classesDir, 0774); err != nil {
		log.Fatalf("failed to run MkdirAll: %+v", err)
	}

	generator := &clientGenerator{
		AppVersion: common.AppVersion,
		OutputDir:  outputFullPath,
	}

	fullPath, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatalf("failed to run filepath.Abs: %+v", err)
	}

	walk(fullPath, "/", generator, &generator.clientType)

	if err := generator.generate(); err != nil {
		log.Fatalf("failed to run generate: %+v", err)
	}
}
