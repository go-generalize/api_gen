package main

import "go/ast"

type PackageStruct struct {
	FileName     string
	PackageName  string
	StructName   string
	StructObject *ast.StructType
}
