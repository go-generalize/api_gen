// Package sample ...
package sample

// PostCreateUserRequest ...
type PostCreateUserRequest struct {
	ID       string `param:"id"`
	Password string
	Gender   int
}

// CreatedType ...
type CreatedType int

// CreatedType Values
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
