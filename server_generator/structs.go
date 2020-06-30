package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
)

func findStructPairList(path string, endpointParams []string) (map[string]*PackageStructPair, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		return nil, err
	}

	structList := findStructList(pkgs)
	structPair := make(map[string]*PackageStructPair, 0)
	for _, s := range structList {
		structEndpointParams := make([]string, len(endpointParams))
		copy(structEndpointParams, endpointParams)

		filePath := fset.File(s.StructObject.Struct).Name()
		fileName := filepath.Base(filePath)
		fileName = fileName[:len(fileName)-len(filepath.Ext(filePath))]

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
		structPair[controllerName].FileName = filePath

		if strings.HasPrefix(fileName, "0_") {
			fileName = fileName[2:]
			if !strings.HasPrefix(strings.ToLower(fileName), ":id") {
				fileName = strings.Replace(fileName, "_id", "ID", -1)
			}
			endpoint := strcase.ToCamel(fileName)
			param := strings.ToLower(string(endpoint[0])) + endpoint[1:]
			structEndpointParams = append(structEndpointParams, param)
		}

		switch structMode {
		case StructModeRequest:
			structPair[controllerName].Request = s
			if len(endpointParams) > 0 {
				structPair[controllerName].LastParam = endpointParams[len(endpointParams)-1]
			}

			if err := validateRequestByEndpointParams(fset, s.StructObject, structEndpointParams); err != nil {
				return nil, err
			}

			if strings.HasPrefix(strings.ToLower(controllerName), "get") {
				if err = validateGetRequestTags(fset, s.StructObject, structEndpointParams); err != nil {
					return nil, err
				}
			}
		case StructModeResponse:
			structPair[controllerName].Response = s
		}
	}

	return structPair, err
}

func validateRequestByEndpointParams(fset *token.FileSet, structType *ast.StructType, endpointParams []string) error {
	fieldList := structType.Fields.List
	hasEndpoints := make(map[string]bool)
	for _, e := range endpointParams {
		hasEndpoints[e] = false
	}

	for _, fields := range fieldList {
		if len(fields.Names) > 1 {
			return fmt.Errorf("%+v: 同じ行に複数のパラメータを記述することはできません。", fields.Names)
		}

		fieldName := fields.Names[0].Name
		if fields.Tag != nil {
			tags := reflect.StructTag(strings.Trim(fields.Tag.Value, "`"))
			if paramTag, ok := tags.Lookup("param"); ok {
				fieldName = paramTag
			}
		}

		if _, ok := hasEndpoints[fieldName]; ok {
			hasEndpoints[fieldName] = true
		}
	}

	requireParams := ""
	for name, e := range hasEndpoints {
		if e {
			continue
		}
		if len(requireParams) > 0 {
			requireParams += ", "
		}
		requireParams += name
	}

	if len(requireParams) == 0 {
		return nil
	}

	return fmt.Errorf("%s: Pathマッチング用のパラメータが不足しています。不足しているパラメータ[%s]",
		fset.Position(structType.Pos()).String(), requireParams)
}

func validateGetRequestTags(fset *token.FileSet, structType *ast.StructType, endpointParams []string) error {
	fieldList := structType.Fields.List
	ep := make(map[string]struct{})
	for _, e := range endpointParams {
		ep[e] = struct{}{}
	}

	for i := range fieldList {
		if fieldList[i].Tag == nil {
			continue
		}

		tags := reflect.StructTag(strings.Trim(fieldList[i].Tag.Value, "`"))

		paramTag, ok := tags.Lookup("param")
		if ok {
			_, ok := ep[paramTag]
			if ok {
				continue
			}
		}

		jsonTag, ok := tags.Lookup("json")

		if !ok {
			continue
		}

		queryTag, ok := tags.Lookup("query")

		if !ok || jsonTag != queryTag {
			return fmt.Errorf(
				"%s: GETのRequest structはjsonタグをqueryタグに同じ値を指定する必要があります",
				fset.Position(fieldList[i].Tag.Pos()).String(),
			)
		}
	}

	return nil
}

var (
	epMap = make(map[string]string)
)

func findStructList(pkgs map[string]*ast.Package) []*PackageStruct {
	structList := make([]*PackageStruct, 0)

	for _, pkg := range pkgs {
		for fileName, f := range pkg.Files {
			objects := f.Scope.Objects
			for _, object := range objects {
				if object.Kind != ast.Typ {
					continue
				}

				name := object.Name
				if strings.HasSuffix(name, "Request") {
					epMap[name] = ""
				}
				tSpec, ok := object.Decl.(*ast.TypeSpec)
				if !ok {
					continue
				}

				t := tSpec.Type
				switch t.(type) {
				case *ast.StructType:
					structObject := t.(*ast.StructType)
					structList = append(structList, &PackageStruct{
						FileName:     fileName,
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
