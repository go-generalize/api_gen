// Package clientts generates client-side TypeScript library
package clientts

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/go-generalize/api_gen/v2/pkg/parser"
	"github.com/go-generalize/api_gen/v2/pkg/util"
	"github.com/iancoleman/strcase"
)

// NewGenerator returns a new client-side TypeScript library generator
func NewGenerator(gr *parser.Group, version string) Generator {
	return &generator{
		AppVersion: version,
		root:       gr,
	}
}

// GenerateClient generates a TypeScript client for api_gen
func (g *generator) GenerateClient() (string, error) {
	g.generateGroup(g.root)
	g.sort()

	tmpl, err := template.ParseFS(clientTSTemplate, "templates/api.ts.tmpl")

	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, g); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g *generator) generateEndpoint(ep *parser.Endpoint) (*endpointType, []importPair) {
	endpoint := &endpointType{
		Method:    strings.ToUpper(string(ep.Method)),
		HasFields: len(ep.ResponsePayload.Fields.List) != 0,
		Name:      strings.ToLower(string(ep.Method)) + strcase.ToCamel(ep.Path),
	}

	endpoint.Endpoint = ep.GetFullPath("/", func(rawPath, path, placeholder string) string {
		if placeholder != "" {
			field := util.FindStructField(ep.RequestGo2tsPayload, parser.QueryParamTag, placeholder)
			jsonKey := field

			for _, f := range ep.RequestGo2tsPayload.Entries {
				if f.RawName == field {
					j, ok := reflect.
						StructTag(f.RawTag).
						Lookup(parser.JSONParamTag)

					if ok {
						jsonKey = j
					}
				}
			}

			return fmt.Sprintf(
				`${encodeURI(param.%s.toString())}`,
				jsonKey,
			)
		}

		return path
	})

	urlParams := make(map[string]string)
	for key, field := range ep.RequestGo2tsPayload.Entries {
		param := field.RawName
		jsonKey := key

		if field.RawTag != "" {
			tags := reflect.StructTag(field.RawTag)

			tag, ok := tags.Lookup(parser.JSONParamTag)

			if ok {
				jsonKey = tag
			}

			tag, ok = tags.Lookup(parser.QueryParamTag)

			if ok {
				param = tag
			}
		}

		urlParams[param] = jsonKey
	}

	endpoint.URLParams = make([]string, 0, 3)
	ep.GetFullPath("", func(rawPath, path, placeholder string) string {
		if placeholder == "" {
			return ""
		}

		endpoint.URLParams = append(endpoint.URLParams, urlParams[placeholder])

		return ""
	})

	sort.Slice(endpoint.URLParams, func(i, j int) bool {
		return endpoint.URLParams[i] < endpoint.URLParams[j]
	})

	typeAliasPrefix := ep.GetParent().GetFullPath("", func(rawPath, path, placeholder string) string {
		return strcase.ToCamel(rawPath)
	})

	endpoint.RequestType = typeAliasPrefix + ep.RequestPayloadName
	endpoint.ResponseType = typeAliasPrefix + ep.ResponsePayloadName

	return endpoint, []importPair{
		{
			Name:   ep.RequestPayloadName,
			NameAs: endpoint.RequestType,
		},
		{
			Name:   ep.ResponsePayloadName,
			NameAs: endpoint.ResponseType,
		},
	}
}

func (g *generator) generateGroup(gr *parser.Group) childrenType {
	client := clientType{}

	className := strcase.ToCamel(gr.GetFullPath("_", func(rawPath, path, placeholder string) string {
		return rawPath
	})) + "Client"

	client.Name = className

	it := importType{
		Path: "./classes" + gr.GetFullPath("/", func(rawPath, path, placeholder string) string {
			return rawPath
		}) + "/types",
	}

	client.Methods = make([]*endpointType, 0, len(gr.Endpoints))
	for _, e := range gr.Endpoints {
		ep, ips := g.generateEndpoint(e)

		client.Methods = append(client.Methods, ep)

		for i := range ips {
			it.Pairs = append(it.Pairs, ips[i])
		}
	}

	client.Children = make([]childrenType, 0, len(gr.Children))
	for _, child := range gr.Children {
		client.Children = append(client.Children, g.generateGroup(child))
	}

	// Update global variables
	g.Imports = append(g.Imports, it)

	if gr.GetParent() == nil {
		g.clientType = client
	} else {
		g.ChildrenClients = append(g.ChildrenClients, &client)
	}

	return childrenType{
		Name:      gr.RawPath,
		ClassName: className,
	}
}

func (g *generator) sort() {
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
