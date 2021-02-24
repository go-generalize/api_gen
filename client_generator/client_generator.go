// Package main ...
package main

import (
	"embed"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/template"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

type methodType struct {
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
	Methods  []*methodType
	Children []childrenType
}

type importPair struct {
	Name, NameAs string
}

type importType struct {
	Path  string
	Pairs []importPair
}

type clientGenerator struct {
	AppVersion string
	Imports    []importType
	clientType
	ChildrenClients []*clientType
	OutputDir       string
}

func (g *clientGenerator) generate() error {
	g.sort()

	f, err := templatesFS.Open("templates/api.ts.tmpl")
	if err != nil {
		return err
	}

	templ, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	t := template.Must(template.New("tmpl").Parse(string(templ)))

	apiTsPath := filepath.Join(g.OutputDir, "./api_client.ts")
	fp, err := os.Create(apiTsPath)

	if err != nil {
		log.Fatalf("failed to create api_client.ts: %+v", err)
	}
	defer fp.Close()

	if err := t.Execute(fp, g); err != nil {
		log.Fatalf("failed to execute template: %+v", err)
	}

	return nil
}

func (g *clientGenerator) sort() {
	sort.Slice(g.Imports, func(i, j int) bool {
		return g.Imports[i].Path < g.Imports[j].Path
	})

	for i := range g.Imports {
		pairs := g.Imports[i].Pairs
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].NameAs < pairs[j].NameAs
		})
	}

	sort.Slice(g.ChildrenClients, func(i, j int) bool {
		return g.ChildrenClients[i].Name < g.ChildrenClients[j].Name
	})

	sort.Slice(g.Methods, func(i, j int) bool {
		return g.Methods[i].Name < g.Methods[j].Name
	})

	for _, clientType := range g.ChildrenClients {
		methods := clientType.Methods
		sort.Slice(methods, func(i, j int) bool {
			return methods[i].Name < methods[j].Name
		})
	}
}
