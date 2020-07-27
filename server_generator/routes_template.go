package main

type RoutesTemplate struct {
	AppVersion  string
	Package     string
	Controllers []*ControllerTemplate
}
