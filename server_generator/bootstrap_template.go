package main

type BootstrapTemplate struct {
	PackageName string
	Bootstraps  []*BootstrapTemplates
}

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
