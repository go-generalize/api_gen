// Package user2 ...
package user2

import "time"

// PostUpdateUserPasswordRequest ...
type PostUpdateUserPasswordRequest struct {
	Password        string
	PasswordConfirm string
}

// PostUpdateUserPasswordResponse ...
type PostUpdateUserPasswordResponse struct {
	Status      bool
	Message     string
	RequestTime time.Time
}
