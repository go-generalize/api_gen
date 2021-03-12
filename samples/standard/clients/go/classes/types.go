// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: devel

package types

import (
	"time"
)

type GetRequest struct {
	Param string
}

type GetResponse struct {
	Data string
}

type Pos struct {
	X int
	Y int
}

type PostCreateTableRequest struct {
	Flag int
	ID   string
	Map  map[int]int `json:"map"`
	Text string
}

type PostCreateTableResponse struct {
	ID          string
	Payload     Table
	RequestTime time.Time
}

type PostCreateUserRequest struct {
	Birthday time.Time
	Gender   int
	ID       string `param:"id"`
	Password string
	Roles    *[]*Role
}

type PostCreateUserResponse struct {
	CreatedType CreatedType
	Message     string
	RequestedAt time.Time
	Status      bool
}

type Role struct {
	ID             int
	Name           string
	RecursionRoles *[]Role
}

type Table struct {
	Pos Pos
}

type CreatedType int

const (
	CreatedTypeGuest  int = 2
	CreatedTypeMember int = 1
	CreatedTypeOwner  int = 0
)
