package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"text/template"

	"github.com/rakyll/statik/fs"
)

type methodType struct {
	Name                      string
	RequestType, ResponseType string
	Method, Endpoint          string
	URLParams                 []string
}

type childrenType struct {
	Name, ClassName string
}

type clientType struct {
	Name     string
	Methods  []methodType
	Children []childrenType
}

type importType struct {
	Path         string
	Name, NameAs string
}

type clientGenerator struct {
	AppVersion string
	Imports    []importType
	clientType
	ChildrenClients []clientType
}

func (g *clientGenerator) generate() error {
	g.sort()

	statikFs, err := fs.New()
	if err != nil {
		return err
	}

	f, err := statikFs.Open("/templates/api.ts.tmpl")
	if err != nil {
		return err
	}

	templ, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	t := template.Must(template.New("tmpl").Parse(string(templ)))

	fp, err := os.Create("api_client.ts")

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
		return g.Imports[i].Name+g.Imports[i].Path < g.Imports[j].Name+g.Imports[j].Path
	})

	sort.Slice(g.ChildrenClients, func(i, j int) bool {
		return g.ChildrenClients[i].Name < g.ChildrenClients[j].Name
	})

	sort.Slice(g.Methods, func(i, j int) bool {
		return g.Methods[i].Name < g.Methods[j].Name
	})

	for index := range g.ChildrenClients {
		sort.Slice(g.ChildrenClients[index].Methods, func(i, j int) bool {
			return g.ChildrenClients[index].Methods[i].Name < g.ChildrenClients[index].Methods[j].Name
		})
	}
}
