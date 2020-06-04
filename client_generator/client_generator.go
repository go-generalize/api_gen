package main

import (
	"log"
	"os"
	"text/template"
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

func (g *clientGenerator) generate() {
	t := template.Must(template.New("").Parse(tmpl))

	fp, err := os.Create("api_client.ts")

	if err != nil {
		log.Fatalf("failed to create api_client.ts: %+v", err)
	}
	defer fp.Close()

	if err := t.Execute(fp, g); err != nil {
		log.Fatalf("failed to execute template: %+v", err)
	}
}

const tmpl = `// THIS CODE WAS GENERATED AUTOMATICALLY
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS

{{range $index, $elem := .Imports}}import {{"{"}} {{$elem.Name}} as {{$elem.NameAs}} {{"}"}} from '{{$elem.Path}}';
export {{"{"}} {{$elem.Name}} as {{$elem.NameAs}} {{"}"}} from '{{$elem.Path}}';
{{end}}

{{range $index, $elem := .ChildrenClients}}
class {{$elem.Name}} {
{{range $index, $elem := .Children}}
	public {{$elem.Name}}: {{$elem.ClassName}};{{end}}
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {
{{range $index, $elem := $elem.Children}}
		this.{{$elem.Name}} = new {{$elem.ClassName}}(headers, options, baseURL);{{end}}
	}
{{range $index, $method := $elem.Methods}}
	async {{$method.Name}}(
		param: {{$method.RequestType}},
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<{{$method.ResponseType}}> {
{{if eq $method.Method "GET"}}		const resp = await fetch(
			this.baseURL + "{{$method.Endpoint}}?" + (new URLSearchParams(param.toObject())).toString(),
			{
				method: "{{$method.Method}}",
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);
{{else}}		const resp = await fetch(
			this.baseURL + "{{$method.Endpoint}}",
			{
				method: "{{$method.Method}}",
				body: JSON.stringify(param),
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);
{{end}}
		return new {{$method.ResponseType}}(await resp.json());
	}{{end}}
}
{{end}}

export class APIClient {
	private headers: {[key: string]: string};
	private options: {[key: string]: any};
	private baseURL: string;
{{range $index, $elem := .Children}}
	public {{$elem.Name}}: {{$elem.ClassName}};{{end}}

	constructor(
		token?: string,
		commonHeaders?: {[key: string]: string},
		baseURL?: string,
		commonOptions: {[key: string]: any} = {}
	) {
		const headers: {[key: string]: string} =  {
			'Content-Type': 'application/json',
			...commonHeaders,
		};

		if(token !== undefined)  {
			headers['Authorization'] = 'Bearer ' + token;
		}
		
		this.baseURL =  (baseURL === undefined) ? "" : baseURL;
		this.options = commonOptions;
		this.headers = headers;

{{range $index, $elem := .Children}}
		this.{{$elem.Name}} = new {{$elem.ClassName}}(headers, this.options, this.baseURL);{{end}}
	}
{{range $index, $method := .Methods}}
	async {{$method.Name}}(
		param: {{$method.RequestType}},
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<{{$method.ResponseType}}> {
{{if eq $method.Method "GET"}}		const resp = await fetch(
			this.baseURL + "{{$method.Endpoint}}?" + (new URLSearchParams(param.toObject())).toString(),
			{
				method: "{{$method.Method}}",
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);
{{else}}		const resp = await fetch(
			this.baseURL + "{{$method.Endpoint}}",
			{
				method: "{{$method.Method}}",
				body: JSON.stringify(param),
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);
{{end}}
		return new {{$method.ResponseType}}(await resp.json());
	}{{end}}
}
`
