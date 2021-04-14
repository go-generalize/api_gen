// Package mock provides embed.FS for mock json
// generated version: devel
package mock

import (
	"embed"
)

// MockJSONFS provides embed.FS for mock json
//go:embed json/*
var MockJSONFS embed.FS

// HeaderOption is the header of Api-Gen-Option in JSON
type HeaderOption struct {
	WaitMS     int64  `json:"wait_ms"`
	TargetFile string `json:"target_file"`
}