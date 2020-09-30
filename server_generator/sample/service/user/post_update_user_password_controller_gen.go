// Package user ...
// generated version: unknown
package user

import (
	props "github.com/go-generalize/api_gen/server_generator/sample/props"
	"github.com/labstack/echo/v4"
)

// PostUpdateUserPasswordController ...
type PostUpdateUserPasswordController struct {
	*props.ControllerProps
}

// NewPostUpdateUserPasswordController ...
func NewPostUpdateUserPasswordController(props *props.ControllerProps) *PostUpdateUserPasswordController {
	p := &PostUpdateUserPasswordController{
		ControllerProps: props,
	}
	return p
}

// PostUpdateUserPassword ...
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
	c echo.Context, req *PostUpdateUserPasswordRequest,
) (res *PostUpdateUserPasswordResponse, err error) {
	panic("require implements.")
}
