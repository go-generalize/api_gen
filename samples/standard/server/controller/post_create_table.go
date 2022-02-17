// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: devel
// nolint:dupl
package controller

import (
	"fmt"

	types "github.com/go-generalize/api_gen/v2/samples/standard/api"
	props "github.com/go-generalize/api_gen/v2/samples/standard/server/props"
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
	if req.File != nil {
		fmt.Println("file", req.File.Filename, req.File.Size)
	}
	if req.Files != nil {
		for _, f := range req.Files {
			fmt.Println("files", f.Filename, f.Size)
		}
	}
	fmt.Println(*req)

	return &types.PostCreateTableResponse{
		ID: "id",
	}, nil

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
	return false
}
