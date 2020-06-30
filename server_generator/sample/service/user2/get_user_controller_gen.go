package user2

import (
	"context"

	"github.com/labstack/echo/v4"
)

type GetUserController struct {
}

func NewGetUserController() *GetUserController {
	g := &GetUserController{}
	return g
}

func (g *GetUserController) GetUser(
	ctx context.Context, c echo.Context, req *GetUserRequest,
) (res *GetUserResponse, err error) {
	return &GetUserResponse{
		ID:            req.ID,
		SearchRequest: req.SearchRequest,
	}, nil
}
