// Package main ...
package main

import (
	"fmt"
	"strings"
)

func getDefaultJsonName(c *ControllerTemplate) string {
	m := strings.ToLower(c.HTTPMethod)
	return fmt.Sprintf("%s-%s", m, c.Endpoint)
}
