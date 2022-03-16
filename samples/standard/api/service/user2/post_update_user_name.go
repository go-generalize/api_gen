// Package user2 ...
package user2

import (
	"mime/multipart"
	"time"
)

// PostUpdateUserNameRequest ...
type PostUpdateUserNameRequest struct {
	Name  string
	File  *multipart.FileHeader   `form:"file"`
	Files []*multipart.FileHeader `form:"files"`
}

// PostUpdateUserNameResponse ...
type PostUpdateUserNameResponse struct {
	Status      bool
	Message     string
	RequestTime time.Time
}
