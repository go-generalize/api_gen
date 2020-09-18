package main

// BootstrapTemplate ...
type BootstrapTemplate struct {
	AppVersion             string
	PackageName            string
	Bootstraps             []*BootstrapTemplates
	ControllerPropsPackage string
}

// BootstrapTemplates ...
type BootstrapTemplates struct {
	ParentPackageName string
	RouteGroupName    string
	HasParent         bool
	PackagePath       string
	ImportPackageName string
	Endpoint          string
	EndpointPath      string
	Controller        *ControllerTemplate
}
