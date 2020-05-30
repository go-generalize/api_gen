package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"regexp"
	"strings"
)

func findStructPairList(path string) (map[string]*PackageStructPair, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.AllErrors|parser.ParseComments)
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

			if strings.HasPrefix(strings.ToLower(controllerName), "get") {
				if err = validateGetRequestTags(fset, s.StructObject); err != nil {
					return nil, err
				}
			}
		case StructModeResponse:
			structPair[controllerName].Response = s
		}
	}

	return structPair, err
}

func validateGetRequestTags(fset *token.FileSet, structType *ast.StructType) error {
	fieldList := structType.Fields.List

	for i := range fieldList {
		if fieldList[i].Tag == nil {
			continue
		}

		tags := reflect.StructTag(strings.Trim(fieldList[i].Tag.Value, "`"))

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
	epMap    = make(map[string]string)
	assigned = regexp.MustCompile(`//\s(.+)ep="(.+)"$`)
)

func findCommentList(commentGroup *ast.CommentGroup) {
	for _, comment := range commentGroup.List {
		group := assigned.FindStringSubmatch(comment.Text)
		if len(group) == 0 {
			continue
		}
		name := strings.Split(group[1], " ")[0]
		if val, ok := epMap[name]; ok && val == "" {
			epMap[name] = group[len(group)-1]
		}
	}
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
			if f.Comments != nil {
				for _, c := range f.Comments {
					findCommentList(c)
				}
			}
		}
	}

	return structList
}
