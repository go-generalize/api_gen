// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

package types

import (
	"time"
)

type Company struct {
	Metadata Metadata
}

type Department struct {
	Metadata Metadata
}

type GetGroupsRequest struct {
}

type GetGroupsResponse struct {
	Companies   *[]Company
	Departments *[]Department
}

type Metadata struct {
	CreatedAt time.Time
	DeletedAt *time.Time
	ID        string
	Name      string
	UpdatedAt time.Time
}
