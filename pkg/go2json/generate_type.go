// Package go2json ...
package go2json

import "github.com/go-generalize/go-easyparser/types"

// MockPayload ...
type MockPayload interface{}

// GenerateType mock json
type GenerateType struct {
	Meta    *MockMeta   `json:"meta"`
	Payload MockPayload `json:"payload"`
}

// MockMeta ...
type MockMeta struct {
	Status       int         `json:"status"`
	MatchRequest interface{} `json:"match_request"`
}

// GenerateSet ...
type GenerateSet struct {
	Request  types.Type
	Response types.Type
}
