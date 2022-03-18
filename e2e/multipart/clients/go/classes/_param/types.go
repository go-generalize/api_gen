// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

package types

import (
	external_3524d37 "github.com/go-generalize/multipart-util"
)

type PostBRequest struct {
	File    *external_3524d37.MultipartPayload   `form:"file" json:"-"`
	Files   []*external_3524d37.MultipartPayload `form:"files" json:"-"`
	Param   string                               `param:"param"`
	Payload string                               `json:"payload"`
}

type PostBResponse struct {
	Message string `json:"message"`
}
