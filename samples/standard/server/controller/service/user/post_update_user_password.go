// Package controller is for
// generated version: devel
// nolint:dupl
package controller

import (
	types "github.com/go-generalize/api_gen/samples/standard/api/service/user"
	props "github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// PostUpdateUserPasswordController ...
type PostUpdateUserPasswordController interface {
	PostUpdateUserPassword(c echo.Context, req *types.PostUpdateUserPasswordRequest) (res *types.PostUpdateUserPasswordResponse, err error)
}

type postUpdateUserPasswordController struct {
	*props.ControllerProps
}

// NewPostUpdateUserPasswordController ...
func NewPostUpdateUserPasswordController(cp *props.ControllerProps) PostUpdateUserPasswordController {
	return &postUpdateUserPasswordController{
		ControllerProps: cp,
	}
}

// PostUpdateUserPassword - POST update_user_password
func (ctrl *postUpdateUserPasswordController) PostUpdateUserPassword(
	c echo.Context, req *types.PostUpdateUserPasswordRequest,
) (res *types.PostUpdateUserPasswordResponse, err error) {
	// return nil, apierror.NewAPIError(http.StatusBadRequest)
	//
	// return nil, apierror.NewAPIError(http.StatusBadRequest).SetError(err)
	//
	// body := map[string]interface{}{
	// 	"code": http.StatusBadRequest,
	// 	"message": "invalid request parameter.",
	// }
	// return nil, apierror.NewAPIError(http.StatusBadRequest, body).SetError(err)
	panic("require implements.") // FIXME require implements.
}

// AutoBind - use echo.Bind
func (ctrl *postUpdateUserPasswordController) AutoBind() bool {
	return true
}
