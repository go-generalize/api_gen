package user2

type PostUpdateUserNameRequest struct {
	Name string
}

type PostUpdateUserNameResponse struct {
	Status  bool
	Message string
}
