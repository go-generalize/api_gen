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

// GetUserController
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} WIP
// @Router /service/user2/{userID} [GET]
func (g *GetUserController) GetUser(
	ctx context.Context, c echo.Context, req *GetUserRequest,
) (res *GetUserResponse, err error) {
	return &GetUserResponse{
		ID:            req.ID,
		SearchRequest: req.SearchRequest,
	}, nil
}
