package user

type PostUpdateUserPasswordRequest struct {
	Password        string
	PasswordConfirm string
}

type PostUpdateUserPasswordResponse struct {
	Status  bool
	Message string
}
