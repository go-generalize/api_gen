// Package main ...
package main

import (
	"fmt"
)

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

func getNewMockRoute(b *BootstrapTemplates) string {
	if b.Controller == nil {
		return ""
	}

	join := fmt.Sprintf(`filepath.Join(jsonDir, "%s")`, b.RawEndpointFilePath)
	fn := fmt.Sprintf(`NewMockRoutes(p, %sGroup, %s, w)`, b.RouteGroupName, join)

	if b.Endpoint == "" {
		return fn
	}
	return fmt.Sprintf("%s.%s", b.RouteGroupName, fn)
}
