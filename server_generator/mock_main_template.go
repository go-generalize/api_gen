// Package main ...
package main

// CreateMockRequest create mock request object
type CreateMockRequest struct {
	RootPath               string
	ControllerPropsPackage string
	ApiRootPackage         string
	BootstrapTemplate      *BootstrapTemplate
	APIRootPathRel         string
}

// MockMainTemplate ...
type MockMainTemplate struct {
	AppVersion             string
	ControllerPropsPackage string
	APIPackageRoot         string
	APIRootPackageName     string
	DefaultJSONDirPath     string
}
