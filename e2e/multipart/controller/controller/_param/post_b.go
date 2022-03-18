// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: (devel)
// nolint:dupl
package controller

import (
	types "github.com/go-gneralize/api_gen/v2/e2e/multipart/api/_param"
	props "github.com/go-gneralize/api_gen/v2/e2e/multipart/controller/props"
	"github.com/labstack/echo/v4"
)

// PostBController ...
type PostBController interface {
	PostB(c echo.Context, req *types.PostBRequest) (res *types.PostBResponse, err error)
}

type postBController struct {
	*props.ControllerProps
}

// NewPostBController ...
func NewPostBController(cp *props.ControllerProps) PostBController {
	return &postBController{
		ControllerProps: cp,
	}
}

// PostB - POST b
func (ctrl *postBController) PostB(
	c echo.Context, req *types.PostBRequest,
) (res *types.PostBResponse, err error) {
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
func (ctrl *postBController) AutoBind() bool {
	return true
}
