// generated version: 0.3.5

package user2

import (
	"context"

	"github.com/labstack/echo/v4"
)

// GetUserController ...
type GetUserController struct {
}

// NewGetUserController ...
func NewGetUserController() *GetUserController {
	g := &GetUserController{}
	return g
}

// GetUser ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param userID path string WIP:${isRequire} WIP:${description}
// @Param search_request query string WIP:${isRequire} WIP:${description}
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
