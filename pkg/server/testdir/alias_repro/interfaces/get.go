// Package interfaces contains request/response types for alias_repro regression test.
// This file intentionally uses map[string]any which triggers *types.Alias in Go 1.23+
// with GODEBUG=gotypesalias=1 (the default since Go 1.23).
package interfaces

// GetRequest - request type with map[string]any (triggers *types.Alias in Go 1.23+)
type GetRequest struct {
	Extra map[string]any `json:"extra"`
}

// GetResponse - response type with map[string]any
type GetResponse struct {
	Data map[string]any `json:"data"`
}
