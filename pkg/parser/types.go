package parser

import "go/ast"

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
