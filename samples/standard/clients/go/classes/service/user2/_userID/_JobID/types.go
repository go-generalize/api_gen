// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: devel

package types

import (
	"time"
)

type PutJobRequest struct {
	JobID  string
	UserID string `param:"userID"`
}

type PutJobResponse struct {
	JobID       string
	RequestTime time.Time
	UserID      string
}
