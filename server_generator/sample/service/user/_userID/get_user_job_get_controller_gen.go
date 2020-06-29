package _userID

import (
	"context"

	"github.com/labstack/echo/v4"
)

type GetUserJobGetController struct {
}

func NewGetUserJobGetController() *GetUserJobGetController {
	g := &GetUserJobGetController{}
	return g
}

func (g *GetUserJobGetController) GetUserJobGet(
	ctx context.Context, c echo.Context, req *GetUserJobGetRequest,
) (res *GetUserJobGetResponse, err error) {
	panic("require implements.")
}
