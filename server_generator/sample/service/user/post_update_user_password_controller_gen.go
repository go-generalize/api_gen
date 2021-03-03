// Package user ...
// generated version: devel
package user

import (
	"io"
	"os"

	"github.com/go-generalize/api_gen/server_generator/sample/apigen/props"
	"github.com/labstack/echo/v4"
)

// PostUpdateUserPasswordController ...
type PostUpdateUserPasswordController struct {
	*props.ControllerProps
}

// NewPostUpdateUserPasswordController ...
func NewPostUpdateUserPasswordController(cp *props.ControllerProps) *PostUpdateUserPasswordController {
	p := &PostUpdateUserPasswordController{
		ControllerProps: cp,
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
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /service/user/update_user_password [POST]
func (p *PostUpdateUserPasswordController) PostUpdateUserPassword(
	c echo.Context, req *PostUpdateUserPasswordRequest,
) (res *PostUpdateUserPasswordResponse, err error) {
	// nolint:errcheck
	io.Copy(os.Stdout, c.Request().Body)

	return &PostUpdateUserPasswordResponse{}, nil
}

// AutoBind - Call c.Bind() before controllers and save the parameters in req
func (p *PostUpdateUserPasswordController) AutoBind() bool {
	return false
}
