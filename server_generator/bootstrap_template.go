package main

type BootstrapTemplate struct {
	PackageName string
	Bootstraps  []*BootstrapTemplates
}

type BootstrapTemplates struct {
	ParentIndex       int
	PackagePath       string
	ImportPackageName string
	Endpoint          string
	EndpointPath      string
	Controller        *ControllerTemplate
}
