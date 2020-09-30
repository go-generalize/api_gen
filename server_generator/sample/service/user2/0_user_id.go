// Package user2 ...
package user2

// GetUserRequest ...
type GetUserRequest struct {
	ID            string `json:"id" param:"userID"`
	SearchRequest string `json:"search_request" query:"search_request"`
}

// GetUserResponse ...
type GetUserResponse struct {
	ID            string
	SearchRequest string
}
