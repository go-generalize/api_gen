// Package main ...
package main

// CreateMockRequest create mock request object
type CreateMockRequest struct {
	RootPath               string
	ControllerPropsPackage string
	ApiRootPackage         string
	BootstrapTemplate      *BootstrapTemplate
	ApiRootPathRel         string
}

// MockMainTemplate ...
type MockMainTemplate struct {
	AppVersion             string
	ControllerPropsPackage string
	ApiPackageRoot         string
	ApiRootPackageName     string
	DefaultJsonDirPath     string
}
