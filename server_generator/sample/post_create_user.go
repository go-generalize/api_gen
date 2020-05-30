package sample

// PostCreateUserRequest Description endpoint"create_user/:id"
type PostCreateUserRequest struct {
	ID       string `json:"id" param:"id"`
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
