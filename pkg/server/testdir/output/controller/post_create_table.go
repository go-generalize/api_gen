// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: devel
// nolint:dupl
package controller

import (
	props "github.com/go-generalize/api_gen/v2/pkg/server/tmp/output/props"
	types "github.com/go-generalize/api_gen/v2/samples/standard/api"
	"github.com/labstack/echo/v4"
)

// PostCreateTableController ...
type PostCreateTableController interface {
	PostCreateTable(c echo.Context, req *types.PostCreateTableRequest) (res *types.PostCreateTableResponse, err error)
}

type postCreateTableController struct {
	*props.ControllerProps
}

// NewPostCreateTableController ...
func NewPostCreateTableController(cp *props.ControllerProps) PostCreateTableController {
	return &postCreateTableController{
		ControllerProps: cp,
	}
}

// PostCreateTable - POST create_table
func (ctrl *postCreateTableController) PostCreateTable(
	c echo.Context, req *types.PostCreateTableRequest,
) (res *types.PostCreateTableResponse, err error) {
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
func (ctrl *postCreateTableController) AutoBind() bool {
	return true
}
