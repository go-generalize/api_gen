package user

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
// @Param Password body string WIP:${isRequire} WIP:${description}
// @Param PasswordConfirm body string WIP:${isRequire} WIP:${description}
// @Success 200 {object} PostUpdateUserPasswordResponse
// @Failure 400 {object} WIP
// @Router /service/user/update_user_password [POST]
func (p *PostUpdateUserPasswordController) PostUpdateUserPassword(
	ctx context.Context, c echo.Context, req *PostUpdateUserPasswordRequest,
) (res *PostUpdateUserPasswordResponse, err error) {
	panic("require implements.")
}
