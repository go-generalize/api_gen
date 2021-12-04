// Package common ...
package common

import "runtime/debug"

// AppVersion is a version number for this module.
var AppVersion = "devel"

func init() {
	if AppVersion != "devel" {
		return
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	AppVersion = bi.Main.Version
}
