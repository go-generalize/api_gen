// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: devel
// nolint:dupl
package controller

import (
	types "github.com/go-generalize/api_gen/v2/samples/standard/api"
	props "github.com/go-generalize/api_gen/v2/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// PostCreateUserController ...
type PostCreateUserController interface {
	PostCreateUser(c echo.Context, req *types.PostCreateUserRequest) (res *types.PostCreateUserResponse, err error)
}

type postCreateUserController struct {
	*props.ControllerProps
}

// NewPostCreateUserController ...
func NewPostCreateUserController(cp *props.ControllerProps) PostCreateUserController {
	return &postCreateUserController{
		ControllerProps: cp,
	}
}

// PostCreateUser - POST create_user
func (ctrl *postCreateUserController) PostCreateUser(
	c echo.Context, req *types.PostCreateUserRequest,
) (res *types.PostCreateUserResponse, err error) {
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
func (ctrl *postCreateUserController) AutoBind() bool {
	return true
}
