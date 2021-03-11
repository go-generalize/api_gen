// Package user2 ...
// generated version: devel
package user2

import (
	"github.com/go-generalize/api_gen/samples/standard/server/props"
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
// @Param Password body string false ""
// @Param PasswordConfirm body string false ""
// @Success 200 {object} PostUpdateUserPasswordResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /service/user2/update_user_password [POST]
func (p *PostUpdateUserPasswordController) PostUpdateUserPassword(
	c echo.Context, req *PostUpdateUserPasswordRequest,
) (res *PostUpdateUserPasswordResponse, err error) {
	// API Error Usage: github.com/go-generalize/api_gen/samples/standard/server/wrapper
	//
	// return nil, wrapper.NewAPIError(http.StatusBadRequest)
	//
	// return nil, wrapper.NewAPIError(http.StatusBadRequest).SetError(err)
	//
	// body := map[string]interface{}{
	// 	"code": http.StatusBadRequest,
	// 	"message": "invalid request parameter.",
	// }
	// return nil, wrapper.NewAPIError(http.StatusBadRequest, body).SetError(err)
	panic("require implements.") // FIXME require implements.
}

// AutoBind - use echo.Bind
func (p *PostUpdateUserPasswordController) AutoBind() bool {
	return true
}
