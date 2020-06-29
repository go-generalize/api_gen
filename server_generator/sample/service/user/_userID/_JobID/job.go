package _JobID

type PutJobRequest struct {
	UserID string `param:"userID"`
	JobID  string
}

type PutJobResponse struct {
}
