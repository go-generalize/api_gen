// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

package types

import (
	"time"
)

type GetRequest struct {
}

type GetResponse struct {
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
