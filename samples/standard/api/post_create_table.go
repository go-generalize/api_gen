// Package sample ...
package sample

import (
	"time"
)

// @Summary Create new table
// @Description Create new table

// Flag - type alias(int)
type Flag int

// PostCreateTableRequest ...
type PostCreateTableRequest struct {
	ID   string
	Text string
	Flag Flag
	Map  map[Flag]Flag `json:"map"`
}

// PostCreateTableResponse ...
type PostCreateTableResponse struct {
	ID          string
	RequestTime time.Time
}
