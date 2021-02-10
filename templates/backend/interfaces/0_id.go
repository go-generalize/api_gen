package interfaces

import "time"

// Example endpoint

// PostCreateUserRequest Description
type PostCreateUserRequest struct {
	ID       string    `param:"id" validate:"required"` // User ID
	Password string    `json:"password"`                // User Password
	Gender   int       `json:"gender"`                  // User Gender
	Birthday time.Time `json:"birthday"`                // User Birthday
}

// CreatedType ...
type CreatedType int

// CreatedType values
const (
	CreatedTypeOwner CreatedType = iota
	CreatedTypeMember
	CreatedTypeGuest
)

// PostCreateUserResponse ...
type PostCreateUserResponse struct {
	Status      bool
	Message     string
	CreatedType CreatedType
}
