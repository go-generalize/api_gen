// Package main ...
package main

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func getJsonDir(c *ControllerTemplate) string {
	m := strings.ToLower(c.HTTPMethod)
	rhn := []rune(c.HandlerName)
	opName := strcase.ToSnake(string(rhn[len(m):]))

	return fmt.Sprintf("%s_%s", m, opName)
}
