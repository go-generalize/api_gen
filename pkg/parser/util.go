package parser

import "path/filepath"

func stem(name string) string {
	base := filepath.Base(name)
	return base[:len(base)-len(filepath.Ext(name))]
}
