package _JobID

import "time"

type PutJobRequest struct {
	UserID string `param:"userID"`
	JobID  string
}

type PutJobResponse struct {
	UserID      string
	JobID       string
	RequestTime time.Time
}
