// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

package types

import (
	external_3524d37 "github.com/go-generalize/multipart-util"
	"time"
)

type DeleteUserRequest struct {
	ID string `json:"id" param:"user_id"`
}

type DeleteUserResponse struct {
}

type Embedding struct {
	Page string `json:"page" query:"p"`
}

type GetUserRequest struct {
	ID            string `json:"id" param:"user_id"`
	Page          string `json:"page" query:"p"`
	SearchRequest string `json:"search_request" query:"search_request"`
}

type GetUserResponse struct {
	ID            string
	RequestTime   time.Time
	SearchRequest string
}

type PostUpdateUserNameRequest struct {
	File  *external_3524d37.MultipartPayload   `form:"file" json:"-"`
	Files []*external_3524d37.MultipartPayload `form:"files" json:"-"`
	Name  string
}

type PostUpdateUserNameResponse struct {
	Message     string
	RequestTime time.Time
	Status      bool
}

type PostUpdateUserPasswordRequest struct {
	Password        string
	PasswordConfirm string
}

type PostUpdateUserPasswordResponse struct {
	Message     string
	RequestTime time.Time
	Status      bool
}
