package sample

import "github.com/go-generalize/api_gen/server_generator/sample/service/table"

type PostCreateTableRequest struct {
	ID   string
	Text string
}

type PostCreateTableResponse struct {
	ID      string
	Payload table.Table
}
