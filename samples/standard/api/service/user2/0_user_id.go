// Package user2 ...
package user2

import "time"

// @Summary Get user information by user id

type Embedding struct {
	Page string `json:"page" query:"p"`
}

// GetUserRequest ...
type GetUserRequest struct {
	ID            string `json:"id" param:"user_id"`
	SearchRequest string `json:"search_request" query:"search_request"`
	Embedding
}

// GetUserResponse ...
type GetUserResponse struct {
	ID            string
	SearchRequest string
	RequestTime   time.Time
}

// @Summary Delete a user by user id

// DeleteUserRequest ...
type DeleteUserRequest struct {
	ID string `json:"id" param:"user_id"`
}

// DeleteUserResponse ...
type DeleteUserResponse struct {
}
