// Package controller is for
// generated version: devel
// nolint:dupl
package controller

import (
	types "github.com/go-generalize/api_gen/samples/standard/api/service/user2"
	props "github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// GetUserController ...
type GetUserController interface {
	GetUser(c echo.Context, req *types.GetUserRequest) (res *types.GetUserResponse, err error)
}

type getUserController struct {
	*props.ControllerProps
}

// NewGetUserController ...
func NewGetUserController(cp *props.ControllerProps) GetUserController {
	return &getUserController{
		ControllerProps: cp,
	}
}

// GetUser - GET user
func (ctrl *getUserController) GetUser(
	c echo.Context, req *types.GetUserRequest,
) (res *types.GetUserResponse, err error) {
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
func (ctrl *getUserController) AutoBind() bool {
	return true
}
