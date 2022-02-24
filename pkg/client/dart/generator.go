// Package clientdart generates client-side Dart library
package clientdart

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/go-generalize/api_gen/v2/pkg/parser"
	"github.com/go-generalize/api_gen/v2/pkg/util"
	go2dartgenerator "github.com/go-generalize/go2dart"
	"github.com/iancoleman/strcase"
)

// NewGenerator returns a new client-side Dart library generator
func NewGenerator(gr *parser.Group, version string) Generator {
	return &generator{
		AppVersion: version,
		root:       gr,
	}
}

// GenerateClient generates a Dart client for api_gen
func (g *generator) GenerateClient() (string, error) {
	g.generateGroup(g.root)
	g.sort()

	tmpl, err := template.New("api.dart").Funcs(
		template.FuncMap{
			"toLower": strings.ToLower,
		},
	).ParseFS(clientdartTemplate, "templates/api.dart")

	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	if err := tmpl.ExecuteTemplate(buf, "api.dart", g); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g *generator) generateEndpoint(ep *parser.Endpoint, packageAlias string) *endpointType {
	endpoint := &endpointType{
		Method:    strings.ToUpper(string(ep.Method)),
		HasFields: len(ep.ResponsePayload.Fields.List) != 0,
		Name:      strings.ToLower(string(ep.Method)) + strcase.ToCamel(ep.Path),
		Multipart: ep.UseMultipartUpload,
	}

	endpoint.Endpoint = ep.GetFullPath("/", func(rawPath, path, placeholder string) string {
		if placeholder != "" {
			field := util.FindStructField(ep.RequestGo2tsPayload, parser.QueryParamTag, placeholder)

			return fmt.Sprintf(
				`${Uri.encodeComponent(param.%s.toString())}`,
				go2dartgenerator.ReplaceFieldName(field),
			)
		}

		return path
	})

	res, _ := parser.GetFileFields(ep.RequestGo2tsPayload)
	fileFields := make([]fileField, 0, len(res))

	for _, v := range res {
		isArray := v.Type == parser.UploadMultipleFiles
		fileFields = append(fileFields, fileField{
			MultipartField: v.FormTag,
			StructField:    go2dartgenerator.ReplaceFieldName(v.Value.RawName),
			IsArray:        isArray,
		})

		if isArray {
			g.ImportHTTPParser = true
		}
	}

	endpoint.FileFields = fileFields

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

	endpoint.RequestType = packageAlias + "." + ep.RequestPayloadName
	endpoint.ResponseType = packageAlias + "." + ep.ResponsePayloadName

	return endpoint
}

func (g *generator) generateGroup(gr *parser.Group) childrenType {
	client := clientType{}

	className := strcase.ToCamel(gr.GetFullPath("_", func(rawPath, path, placeholder string) string {
		return rawPath
	})) + "Client"

	client.Name = className

	packageAlias := "types_" + gr.GetFullPath("_", func(rawPath, path, placeholder string) string {
		return strcase.ToSnake(rawPath)
	})
	it := importType{
		Path: "./classes" + gr.GetFullPath("/", func(rawPath, path, placeholder string) string {
			return rawPath
		}) + "/types",
		Alias: packageAlias,
	}

	client.Methods = make([]*endpointType, 0, len(gr.Endpoints))
	for _, e := range gr.Endpoints {
		ep := g.generateEndpoint(e, packageAlias)

		client.Methods = append(client.Methods, ep)
	}

	client.Children = make([]childrenType, 0, len(gr.Children))
	for _, child := range gr.Children {
		client.Children = append(client.Children, g.generateGroup(child))
	}

	if len(client.Methods) != 0 {
		// Update global variables
		g.Imports = append(g.Imports, it)
	}

	if gr.GetParent() == nil {
		g.clientType = client
	} else {
		g.ChildrenClients = append(g.ChildrenClients, &client)
	}

	return childrenType{
		Name:      strings.TrimPrefix(gr.RawPath, "_"),
		ClassName: className,
	}
}

func (g *generator) sort() {
	sort.Slice(g.Imports, func(i, j int) bool {
		return g.Imports[i].Path < g.Imports[j].Path
	})

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
