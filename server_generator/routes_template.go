// Package main ...
package main

// RoutesTemplate ...
type RoutesTemplate struct {
	AppVersion             string
	Package                string
	Controllers            []*ControllerTemplate
	ControllerPropsPackage string
	WrapperPackage         string
}
