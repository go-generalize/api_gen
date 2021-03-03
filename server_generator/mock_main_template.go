// Package main ...
package main

// CreateMockRequest create mock request object
type CreateMockRequest struct {
	RootPath               string
	BootstrapPackage       string
	ControllerPropsPackage string
	APIRootPackage         string
	BootstrapTemplate      *BootstrapTemplate
	APIRootPathRel         string
}

// MockMainTemplate ...
type MockMainTemplate struct {
	AppVersion             string
	ControllerPropsPackage string
	BootstrapPackage       string
	DefaultJSONDirPath     string
}
