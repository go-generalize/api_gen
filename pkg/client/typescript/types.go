// Package clientts generates client-side TypeScript library
package clientts

import (
	"embed"
)

//go:embed templates/*.tmpl
var clientTSTemplate embed.FS

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
	AppVersion string
	Imports    []importType
	clientType
	ChildrenClients []*clientType
	OutputDir       string
}
