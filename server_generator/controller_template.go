package main

// ControllerTemplate ...
type ControllerTemplate struct {
	AppVersion            string
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
