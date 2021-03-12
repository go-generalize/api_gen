// Package user ...
package user

import "time"

// PostUpdateUserNameRequest ...
type PostUpdateUserNameRequest struct {
	Name string
}

// PostUpdateUserNameResponse ...
type PostUpdateUserNameResponse struct {
	Status      bool
	Message     string
	RequestTime time.Time
}
