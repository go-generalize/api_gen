package user2

// PostUpdateUserPasswordRequest ...
type PostUpdateUserPasswordRequest struct {
	Password        string
	PasswordConfirm string
}

// PostUpdateUserPasswordResponse ...
type PostUpdateUserPasswordResponse struct {
	Status  bool
	Message string
}
