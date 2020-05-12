package user

type PostUpdateUserNameRequest struct {
	Name string
}

type PostUpdateUserNameResponse struct {
	Status  bool
	Message string
}
