package sample

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

type PostCreateUserController struct {
}

func NewPostCreateUserController() *PostCreateUserController {
	p := &PostCreateUserController{}
	return p
}

func (p *PostCreateUserController) PostCreateUser(ctx context.Context,
	c echo.Context,
	req *PostCreateUserRequest) (
	res *PostCreateUserResponse,
	err error) {

	return &PostCreateUserResponse{
		Status:      true,
		Message:     fmt.Sprintf("password is %s", req.Password),
		CreatedType: CreatedTypeMember,
	}, nil
}
