// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: devel
// nolint:dupl
package controller

import (
	props "github.com/go-generalize/api_gen/v2/pkg/server/tmp/output/props"
	types "github.com/go-generalize/api_gen/v2/samples/standard/api/service/user2/_userID"
	"github.com/labstack/echo/v4"
)

// GetUserJobGetController ...
type GetUserJobGetController interface {
	GetUserJobGet(c echo.Context, req *types.GetUserJobGetRequest) (res *types.GetUserJobGetResponse, err error)
}

type getUserJobGetController struct {
	*props.ControllerProps
}

// NewGetUserJobGetController ...
func NewGetUserJobGetController(cp *props.ControllerProps) GetUserJobGetController {
	return &getUserJobGetController{
		ControllerProps: cp,
	}
}

// GetUserJobGet - GET user_job_get
func (ctrl *getUserJobGetController) GetUserJobGet(
	c echo.Context, req *types.GetUserJobGetRequest,
) (res *types.GetUserJobGetResponse, err error) {
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
func (ctrl *getUserJobGetController) AutoBind() bool {
	return true
}
