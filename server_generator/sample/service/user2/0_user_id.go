package user2

type GetUserRequest struct {
	ID            string `json:"id" param:"userID"`
	SearchRequest string `json:"search_request" query:"search_request"`
}

type GetUserResponse struct {
	ID            string
	SearchRequest string
}
