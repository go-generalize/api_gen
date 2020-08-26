package sample

import "github.com/go-generalize/api_gen/server_generator/sample/service/table"

// PostCreateTableRequest ...
type PostCreateTableRequest struct {
	ID   string
	Text string
}

// PostCreateTableResponse ...
type PostCreateTableResponse struct {
	ID      string
	Payload table.Table
}
