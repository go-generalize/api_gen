package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func findStructPairList(path string) (map[string]*PackageStructPair, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	structList := findStructList(pkgs)
	structPair := make(map[string]*PackageStructPair, 0)
	for _, s := range structList {
		controllerName := s.StructName
		var structMode StructMode

		if strings.HasSuffix(controllerName, "Request") {
			controllerName = controllerName[:len(controllerName)-7]
			structMode = StructModeRequest
		} else if strings.HasSuffix(controllerName, "Response") {
			controllerName = controllerName[:len(controllerName)-8]
			structMode = StructModeResponse
		} else {
			continue
		}

		if _, ok := structPair[controllerName]; !ok {
			structPair[controllerName] = new(PackageStructPair)
		}

		switch structMode {
		case StructModeRequest:
			structPair[controllerName].Request = s
		case StructModeResponse:
			structPair[controllerName].Response = s
		}
	}

	return structPair, err
}

func findStructList(pkgs map[string]*ast.Package) []*PackageStruct {
	structList := make([]*PackageStruct, 0)

	for _, pkg := range pkgs {
		for _, f := range pkg.Files {
			objects := f.Scope.Objects
			for _, object := range objects {
				if object.Kind != ast.Typ {
					continue
				}

				name := object.Name
				tSpec, ok := object.Decl.(*ast.TypeSpec)
				if !ok {
					continue
				}

				t := tSpec.Type
				switch t.(type) {
				case *ast.StructType:
					structObject := t.(*ast.StructType)
					structList = append(structList, &PackageStruct{
						PackageName:  pkg.Name,
						StructName:   name,
						StructObject: structObject,
					})
				case *ast.Ident:
					//log.Printf("<IDENT> %s (%s)", name, t.(*ast.Ident).Name)
				default:
					//log.Printf("name=%s , %#v\n", name, tSpec)
				}
			}
		}
	}

	return structList
}
