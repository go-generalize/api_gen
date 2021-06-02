// Package user2 ...
package user2

import "time"

// @Summary Get user information by user id

// GetUserRequest ...
type GetUserRequest struct {
	ID            string `json:"id" param:"id"`
	SearchRequest string `json:"search_request" query:"search_request"`
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
