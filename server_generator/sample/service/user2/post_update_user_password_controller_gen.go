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

// PostUpdateUserPassword
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param
// @Success 200 {object} PostUpdateUserPasswordResponse
// @Failure 400 {object} WIP
// @Router /service/user2/update_user_password [POST]
func (p *PostUpdateUserPasswordController) PostUpdateUserPassword(
	ctx context.Context, c echo.Context, req *PostUpdateUserPasswordRequest,
) (res *PostUpdateUserPasswordResponse, err error) {
	panic("require implements.")
}
