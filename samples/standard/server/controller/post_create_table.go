// Package controller is for
// generated version: devel
package controller

import (
	types "github.com/go-generalize/api_gen/samples/standard/api"
	props "github.com/go-generalize/api_gen/samples/standard/server/props"
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
