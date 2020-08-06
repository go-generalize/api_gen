package user

// PostUpdateUserNameRequest ...
type PostUpdateUserNameRequest struct {
	Name string
}

// PostUpdateUserNameResponse ...
type PostUpdateUserNameResponse struct {
	Status  bool
	Message string
}
