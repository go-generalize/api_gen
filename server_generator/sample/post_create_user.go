package sample

// PostCreateUserRequest Description endpoint"user/:userID/create"
type PostCreateUserRequest struct {
	ID       string `json:"userID" param:"userID"`
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
