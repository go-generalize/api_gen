package user2

import (
	"context"

	"github.com/labstack/echo/v4"
)

type PostUpdateUserPasswordController struct {
}

func NewPostUpdateUserPasswordController() *PostUpdateUserPasswordController {
	p := &PostUpdateUserPasswordController{}
	return p
}

func (p *PostUpdateUserPasswordController) PostUpdateUserPassword(ctx context.Context,
	c echo.Context,
	req *PostUpdateUserPasswordRequest) (
	res *PostUpdateUserPasswordResponse,
	err error) {
	panic("require implements.")
}
