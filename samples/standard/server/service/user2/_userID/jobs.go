package _userID

import "time"

type GetUserJobGetRequest struct {
	UserID string `param:"userID"`
}

type GetUserJobGetResponse struct {
	JobName     string
	RequestTime time.Time
}
