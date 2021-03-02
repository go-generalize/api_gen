// Package main ...
package main

import (
	"fmt"
	"strings"
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

	prefix := b.RouteGroupName
	if b.Endpoint == "/" {
		prefix = getEndOfPackage(b.PackagePath)
	}

	return fmt.Sprintf("%s.NewRoutes(p, %sGroup, opts...)", prefix, b.RouteGroupName)
}

func getNewMockRoute(b *BootstrapTemplates) string {
	if b.Controller == nil {
		return ""
	}

	join := fmt.Sprintf(`filepath.Join(jsonDir, "%s")`, b.RawEndpointFilePath)
	fn := fmt.Sprintf(`NewMockRoutes(%sGroup, %s, w)`, b.RouteGroupName, join)

	prefix := b.RouteGroupName
	if b.Endpoint == "/" {
		prefix = getEndOfPackage(b.PackagePath)
	}

	return fmt.Sprintf("%s.%s", prefix, fn)
}

func getEndOfPackage(p string) string {
	sp := strings.Split(p, "/")
	return sp[len(sp)-1]
}
