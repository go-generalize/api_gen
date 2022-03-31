// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: devel
// nolint:dupl
package controller

import (
	props "github.com/go-generalize/api_gen/v2/pkg/server/tmp/output/props"
	types "github.com/go-generalize/api_gen/v2/samples/standard/api/service/groups"
	"github.com/labstack/echo/v4"
)

// GetGroupsController ...
type GetGroupsController interface {
	GetGroups(c echo.Context, req *types.GetGroupsRequest) (res *types.GetGroupsResponse, err error)
}

type getGroupsController struct {
	*props.ControllerProps
}

// NewGetGroupsController ...
func NewGetGroupsController(cp *props.ControllerProps) GetGroupsController {
	return &getGroupsController{
		ControllerProps: cp,
	}
}

// GetGroups - GET groups
func (ctrl *getGroupsController) GetGroups(
	c echo.Context, req *types.GetGroupsRequest,
) (res *types.GetGroupsResponse, err error) {
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
func (ctrl *getGroupsController) AutoBind() bool {
	return true
}
