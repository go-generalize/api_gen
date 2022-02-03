// Package clientdart generates client-side Dart library
package clientdart

import (
	"embed"

	"github.com/go-generalize/api_gen/v2/pkg/parser"
)

//go:embed templates/*.dart
var clientdartTemplate embed.FS

const headerComment = `// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: %s

`

// Generator generates a Dart client for api_gen
type Generator interface {
	GenerateClient() (string, error)
	GenerateTypes(fn func(relPath, code string) error) error
}

type endpointType struct {
	Name                      string
	RequestType, ResponseType string
	Method, Endpoint          string
	URLParams                 []string
	HasFields                 bool
	Multipart                 bool
	FileFieldNames            []string
}

type childrenType struct {
	Name, ClassName string
}

type clientType struct {
	Name     string
	Methods  []*endpointType
	Children []childrenType
}

type importType struct {
	Path  string
	Alias string
}

type generator struct {
	root *parser.Group

	AppVersion string
	Imports    []importType
	clientType
	ChildrenClients []*clientType
	OutputDir       string
}

var _ Generator = &generator{}
