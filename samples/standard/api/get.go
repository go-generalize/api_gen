// Package sample ...
package sample

import "time"

// GetRequest ...
type GetRequest struct {
	Param string
	Time  time.Time
}

// GetResponse ...
type GetResponse struct {
	Data string
}
