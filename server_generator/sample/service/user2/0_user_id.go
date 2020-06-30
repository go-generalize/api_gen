package user2

type GetUserRequest struct {
	ID string `json:"id" param:"userID"`
}

type GetUserResponse struct {
}
