package main

import "go/ast"

type PackageStruct struct {
	PackageName  string
	StructName   string
	StructObject *ast.StructType
}
