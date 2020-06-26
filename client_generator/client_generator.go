package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/rakyll/statik/fs"
)

type methodType struct {
	Name                      string
	RequestType, ResponseType string
	Method, Endpoint          string
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
	Imports []importType
	clientType
	ChildrenClients []clientType
}

func (g *clientGenerator) generate() error {
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
