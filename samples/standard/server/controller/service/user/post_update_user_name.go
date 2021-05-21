// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: devel
// nolint:dupl
package controller

import (
	types "github.com/go-generalize/api_gen/samples/standard/api/service/user"
	props "github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// PostUpdateUserNameController ...
type PostUpdateUserNameController interface {
	PostUpdateUserName(c echo.Context, req *types.PostUpdateUserNameRequest) (res *types.PostUpdateUserNameResponse, err error)
}

type postUpdateUserNameController struct {
	*props.ControllerProps
}

// NewPostUpdateUserNameController ...
func NewPostUpdateUserNameController(cp *props.ControllerProps) PostUpdateUserNameController {
	return &postUpdateUserNameController{
		ControllerProps: cp,
	}
}

// PostUpdateUserName - POST update_user_name
func (ctrl *postUpdateUserNameController) PostUpdateUserName(
	c echo.Context, req *types.PostUpdateUserNameRequest,
) (res *types.PostUpdateUserNameResponse, err error) {
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
func (ctrl *postUpdateUserNameController) AutoBind() bool {
	return true
}
