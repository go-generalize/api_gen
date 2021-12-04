// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

package types

import (
	"time"
)

type GetUserJobGetRequest struct {
	UserID string `param:"userID"`
}

type GetUserJobGetResponse struct {
	JobName     string
	RequestTime time.Time
}
