// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: devel
// nolint:dupl
package controller

import (
	types "github.com/go-generalize/api_gen/samples/standard/api/service/user"
	props "github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// GetController ...
type GetController interface {
	Get(c echo.Context, req *types.GetRequest) (res *types.GetResponse, err error)
}

type getController struct {
	*props.ControllerProps
}

// NewGetController ...
func NewGetController(cp *props.ControllerProps) GetController {
	return &getController{
		ControllerProps: cp,
	}
}

// Get - GET
func (ctrl *getController) Get(
	c echo.Context, req *types.GetRequest,
) (res *types.GetResponse, err error) {
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
func (ctrl *getController) AutoBind() bool {
	return true
}
