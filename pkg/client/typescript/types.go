// Package clientts generates client-side TypeScript library
package clientts

import (
	"embed"

	"github.com/go-generalize/api_gen/pkg/parser"
)

//go:embed templates/*.tmpl
var clientTSTemplate embed.FS

const headerComment = `// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: %s

`

// Generator generates a TypeScript client for api_gen
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
}

type childrenType struct {
	Name, ClassName string
}

type clientType struct {
	Name     string
	Methods  []*endpointType
	Children []childrenType
}

type importPair struct {
	Name, NameAs string
}

type importType struct {
	Path  string
	Pairs []importPair
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
