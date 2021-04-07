// Package sample ...
package sample

import "time"

// Role ...
type Role struct {
	ID             int
	Name           string
	RecursionRoles []Role
}

// PostCreateUserRequest ...
type PostCreateUserRequest struct {
	ID       string `param:"id"`
	Password string
	Gender   int
	Birthday time.Time
	Roles    []*Role
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
	RequestedAt time.Time
}
