package api

import "mime/multipart"

type PostARequest struct {
	File    *multipart.FileHeader   `form:"file"`
	Files   []*multipart.FileHeader `form:"files"`
	Payload string
}

type PostAResponse struct {
	Message string
}
