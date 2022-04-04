// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: devel
// nolint:dupl
package controller

import (
	props "github.com/go-generalize/api_gen/v2/pkg/server/tmp/output/props"
	types "github.com/go-generalize/api_gen/v2/samples/standard/api/service/user2"
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
