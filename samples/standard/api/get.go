// Package sample ...
package sample

import "time"

// Enum ...
type Enum string

const (
	// EnumA ...
	EnumA Enum = "A"
	// EnumB ...
	EnumB Enum = "B"
	// EnumC ...
	EnumC Enum = "C"
)

// GetRequest ...
type GetRequest struct {
	Param string `example:"param param param"`
	Time  time.Time
	Enum  Enum
}

// GetResponse ...
type GetResponse struct {
	Data string
}
