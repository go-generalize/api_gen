package common

import (
	parser "github.com/go-generalize/go-easyparser"
)

// ParserFilter is a common filter of go2ts parser for api_gen
func ParserFilter(opt *parser.FilterOpt) bool {
	if opt.Dependency {
		return true
	}
	if !opt.BasePackage {
		return false
	}
	if !opt.Exported {
		return false
	}

	return true
}
