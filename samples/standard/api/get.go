// Package sample ...
package sample

import "time"

type Enum string

const (
	Enum_A Enum = "A"
	Enum_B Enum = "B"
	Enum_C Enum = "C"
)

// GetRequest ...
type GetRequest struct {
	Param string
	Time  time.Time
	Enum  Enum
}

// GetResponse ...
type GetResponse struct {
	Data string
}
