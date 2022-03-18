package api

import "mime/multipart"

type PostBRequest struct {
	File    *multipart.FileHeader   `form:"file"`
	Files   []*multipart.FileHeader `form:"files"`
	Param   string                  `param:"param"`
	Payload string                  `json:"payload"`
}

type PostBResponse struct {
}
