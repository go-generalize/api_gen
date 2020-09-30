// Package main ...
package main

import "fmt"

func getGroupName(b *BootstrapTemplates) string {
	if !b.HasParent {
		return "e"
	}

	return fmt.Sprintf("%sGroup", b.ParentPackageName)
}

func getNewRoute(b *BootstrapTemplates) string {
	if b.Controller == nil {
		return ""
	}

	if b.Endpoint == "" {
		return fmt.Sprintf("NewRoutes(p, %sGroup)", b.RouteGroupName)
	}

	return fmt.Sprintf("%s.NewRoutes(p, %sGroup)", b.RouteGroupName, b.RouteGroupName)
}
