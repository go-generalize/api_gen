package common

import (
	"strings"

	go2tsparser "github.com/go-generalize/go2ts/pkg/parser"
)

// ParserFilter is a common filter of go2ts parser for api_gen
func ParserFilter(opt *go2tsparser.FilterOpt) bool {
	if opt.Dependency {
		return true
	}
	if !opt.BasePackage {
		return false
	}
	if !opt.Exported {
		return false
	}

	return strings.HasSuffix(opt.Name, "Request") || strings.HasSuffix(opt.Name, "Response")
}
