// Package controller is for
// generated version: devel
package controller

import (
	types "github.com/go-generalize/api_gen/samples/empty_root/api/foo/bar"
	props "github.com/go-generalize/api_gen/samples/empty_root/server/props"
	"github.com/labstack/echo/v4"
)

// PostUserController ...
type PostUserController interface {
	PostUser(c echo.Context, req *types.PostUserRequest) (res *types.PostUserResponse, err error)
}

type postUserController struct {
	*props.ControllerProps
}

// NewPostUserController ...
func NewPostUserController(cp *props.ControllerProps) PostUserController {
	return &postUserController{
		ControllerProps: cp,
	}
}

// PostUser - POST user
func (ctrl *postUserController) PostUser(
	c echo.Context, req *types.PostUserRequest,
) (res *types.PostUserResponse, err error) {
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
func (ctrl *postUserController) AutoBind() bool {
	return true
}
