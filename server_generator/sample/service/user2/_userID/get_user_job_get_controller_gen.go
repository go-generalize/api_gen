// generated version: 0.3.5

package _userID

import (
	"context"

	"github.com/labstack/echo/v4"
)

// GetUserJobGetController ...
type GetUserJobGetController struct {
}

// NewGetUserJobGetController ...
func NewGetUserJobGetController() *GetUserJobGetController {
	g := &GetUserJobGetController{}
	return g
}

// GetUserJobGet ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string WIP:${isRequire} WIP:${description}
// @Success 200 {object} GetUserJobGetResponse
// @Failure 400 {object} WIP
// @Router /service/user2/{userID}/user_job_get [GET]
func (g *GetUserJobGetController) GetUserJobGet(
	ctx context.Context, c echo.Context, req *GetUserJobGetRequest,
) (res *GetUserJobGetResponse, err error) {
	panic("require implements.")
}
