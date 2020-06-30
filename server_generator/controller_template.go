package main

type ControllerTemplate struct {
	Package               string
	ControllerName        string
	ControllerNameInitial string
	HandlerName           string
	RawEndpointPath       string
	Endpoint              string
	HTTPMethod            string
	RequestStructName     string
	ResponseStructName    string
	RequestParams         []RequestParam
}
