package parser

import (
	"go/ast"
	"strings"
)

// MethodType represents methods for HTTP
type MethodType string

const (
	// GET is GET method for HTTP
	GET MethodType = "GET"
	// POST is POST method for HTTP
	POST MethodType = "POST"
	// PUT is PUT method for HTTP
	PUT MethodType = "PUT"
	// DELETE is method for HTTP
	DELETE MethodType = "DELETE"
	// PATCH is method for HTTP
	PATCH MethodType = "PATCH"
)

func getMethodType(structName string) MethodType {
	var method MethodType
	switch lowered := strings.ToLower(structName); {
	case strings.HasPrefix(lowered, string(GET)):
		method = GET
	case strings.HasPrefix(lowered, string(POST)):
		method = POST
	case strings.HasPrefix(lowered, string(PUT)):
		method = PUT
	case strings.HasPrefix(lowered, string(DELETE)):
		method = DELETE
	case strings.HasPrefix(lowered, string(PATCH)):
		method = PATCH
	}

	return method
}

// Endpoint represents one HTTP endpoint
type Endpoint struct {
	parentGroup *Group

	Method MethodType
	Path   string

	RequestPayload  *ast.StructType
	ResponsePayload *ast.StructType
}

// Group is a layer for endpoints
type Group struct {
	ImportPath string
	Path       string
	Dir        string

	Children  []*Group
	Endpoints []*Endpoint
}
