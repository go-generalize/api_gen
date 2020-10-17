// Package main ...
package main

import (
	"fmt"
	"strings"
)

func getDefaultJsonName(c *ControllerTemplate) string {
	m := strings.ToUpper(c.HTTPMethod)
	rhn := []rune(c.HandlerName)
	opName := string(rhn[len(m):])

	return fmt.Sprintf("%s_%s", m, opName)
}
