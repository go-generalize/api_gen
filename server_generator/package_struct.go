package main

import "go/ast"

// PackageStruct ...
type PackageStruct struct {
	FileName      string
	PackageName   string
	StructName    string
	StructObject  *ast.StructType
	RequestParams []RequestParam
}
