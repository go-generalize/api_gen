package user

import (
	"context"

	"github.com/labstack/echo/v4"
)

type PostUpdateUserNameController struct {
}

func NewPostUpdateUserNameController() *PostUpdateUserNameController {
	p := &PostUpdateUserNameController{}
	return p
}

func (p *PostUpdateUserNameController) PostUpdateUserName(ctx context.Context,
	c echo.Context,
	req *PostUpdateUserNameRequest) (
	res *PostUpdateUserNameResponse,
	err error) {
	panic("require implements.")
}
