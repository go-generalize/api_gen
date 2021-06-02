// Package parser parses API definitions in Go
package parser

import (
	"go/ast"
	"go/token"
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

const (
	// QueryParamTag is the tag name for parameters in query
	QueryParamTag = "param"

	// JSONParamTag is the tag name for JSON
	JSONParamTag = "json"
)

var (
	methods = []MethodType{
		GET,
		POST,
		PUT,
		DELETE,
		PATCH,
	}
)

func getMethodType(structName string) MethodType {
	structName = strings.ToUpper(structName)
	for i := range methods {
		if strings.HasPrefix(structName, string(methods[i])) {
			return methods[i]
		}
	}

	return ""
}

// Endpoint represents one HTTP endpoint
type Endpoint struct {
	parentGroup *Group

	Method        MethodType
	RawPath, Path string
	Placeholder   string

	RequestPayloadName  string
	ResponsePayloadName string
	RequestPayload      *ast.StructType
	ResponsePayload     *ast.StructType

	SwagComments []string

	requestPos  token.Pos
	responsePos token.Pos
}

// GetParent returns the parent Group
func (e *Endpoint) GetParent() *Group {
	return e.parentGroup
}

// GetFullPath returns the full path foe the endpoint
func (e *Endpoint) GetFullPath(splitter string, fn func(rawPath, path, placeholder string) string) string {
	return e.parentGroup.GetFullPath(splitter, fn) + splitter + fn(e.RawPath, e.Path, e.Placeholder)
}

// Group is a layer for endpoints
type Group struct {
	parentGroup *Group

	ImportPath    string
	RawPath, Path string
	Dir           string
	Placeholder   string

	Children  []*Group
	Endpoints []*Endpoint
}

// GetParent returns the parent Group
func (g *Group) GetParent() *Group {
	return g.parentGroup
}

// GetFullPath returns the full path foe the endpoint
func (g *Group) GetFullPath(splitter string, fn func(rawPath, path, placeholder string) string) string {
	paths := make([]string, 0, 8)

	paths = append(paths, fn(g.RawPath, g.Path, g.Placeholder))

	gr := g.parentGroup
	for gr != nil {
		paths = append(paths, fn(gr.RawPath, gr.Path, gr.Placeholder))

		gr = gr.parentGroup
	}

	for i := 0; i < len(paths)/2; i++ {
		paths[i], paths[len(paths)-i-1] = paths[len(paths)-i-1], paths[i]
	}

	return strings.Join(paths, splitter)
}
