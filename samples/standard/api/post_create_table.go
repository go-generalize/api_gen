// Package sample ...
package sample

import (
	"mime/multipart"
	"time"
)

// Flag - type alias(int)
type Flag int

// PostCreateTableRequest ...
type PostCreateTableRequest struct {
	ID    string
	Text  string
	Flag  Flag
	Map   map[Flag]Flag           `json:"map"`
	File  *multipart.FileHeader   `form:"file"`
	Files []*multipart.FileHeader `form:"files"`
}

// PostCreateTableResponse ...
type PostCreateTableResponse struct {
	ID          string
	RequestTime time.Time
}
