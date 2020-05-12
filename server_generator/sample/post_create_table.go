package sample

import "github.com/gcp-kit/api_gen/server_generator/sample/service/table"

type PostCreateTableRequest struct {
	ID   string
	Text string
}

type PostCreateTableResponse struct {
	ID      string
	Payload table.Table
}
