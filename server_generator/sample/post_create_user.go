package sample

type PostCreateUserRequest struct {
	ID       string
	Password string
	Gender   int
}

type CreatedType int

const (
	CreatedTypeOwner CreatedType = iota
	CreatedTypeMember
	CreatedTypeGuest
)

type PostCreateUserResponse struct {
	Status      bool
	Message     string
	CreatedType CreatedType
}
