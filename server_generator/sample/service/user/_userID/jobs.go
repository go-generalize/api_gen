package _userID

type GetUserJobGetRequest struct {
	UserID string `param:"userID"`
}

type GetUserJobGetResponse struct {
	JobName string
}
