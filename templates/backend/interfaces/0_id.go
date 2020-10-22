package interfaces

import "time"

// Example endpoint

// PostCreateUserRequest Description
type PostCreateUserRequest struct {
	ID       string `param:"id"`
	Password string
	Gender   int
	Birthday time.Time
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
