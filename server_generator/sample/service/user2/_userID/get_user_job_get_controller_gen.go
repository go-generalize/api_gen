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

// GetUserJobGetController
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param
// @Success 200 {object} GetUserJobGetResponse
// @Failure 400 {object} WIP
// @Router /service/user2/{userID}/user_job_get [GET]
func (g *GetUserJobGetController) GetUserJobGet(
	ctx context.Context, c echo.Context, req *GetUserJobGetRequest,
) (res *GetUserJobGetResponse, err error) {
	panic("require implements.")
}
