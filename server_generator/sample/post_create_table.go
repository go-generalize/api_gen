// Package sample ...
package sample

import (
	"time"

	"github.com/go-generalize/api_gen/server_generator/sample/service/table"
)

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
	Payload     table.Table
	RequestTime time.Time
}
