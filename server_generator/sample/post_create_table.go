// Package sample ...
package sample

import (
	"time"

	"github.com/go-generalize/api_gen/server_generator/sample/service/table"
)

// PostCreateTableRequest ...
type PostCreateTableRequest struct {
	ID   string
	Text string
}

// PostCreateTableResponse ...
type PostCreateTableResponse struct {
	ID          string
	Payload     table.Table
	RequestTime time.Time
}
