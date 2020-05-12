package user2

import (
	"github.com/labstack/echo/v4"
)

type PostUpdateUserNameController struct {
}

func NewPostUpdateUserNameController() *PostUpdateUserNameController {
	p := &PostUpdateUserNameController{}
	return p
}

func (p *PostUpdateUserNameController) PostUpdateUserName(c echo.Context,
	req *PostUpdateUserNameRequest) (
	res *PostUpdateUserNameResponse,
	err error) {
	panic("require implements.")
}
