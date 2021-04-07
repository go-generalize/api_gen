// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: devel

package types

import (
	"time"
)

type GetUserRequest struct {
	ID            string `json:"id" param:"userID"`
	SearchRequest string `json:"search_request" query:"search_request"`
}

type GetUserResponse struct {
	ID            string
	RequestTime   time.Time
	SearchRequest string
}

type PostUpdateUserNameRequest struct {
	Name string
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
