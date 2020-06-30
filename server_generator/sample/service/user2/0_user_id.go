package user2

type GetUserRequest struct {
	ID string `json:"id" param:"userID" query:"id"`
}

type GetUserResponse struct {
}
