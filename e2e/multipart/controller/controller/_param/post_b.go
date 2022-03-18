// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: (devel)
// nolint:dupl
package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-gneralize/api_gen/v2/e2e/e2eutil"
	types "github.com/go-gneralize/api_gen/v2/e2e/multipart/api/_param"
	"github.com/go-gneralize/api_gen/v2/e2e/multipart/controller/pkg/apierror"
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
	{
		file, err := e2eutil.ReadMultipartFile(req.File)

		if err != nil {
			return nil, apierror.NewAPIError(http.StatusBadRequest, err)
		}

		if string(file) != "1" {
			return nil, apierror.NewAPIError(http.StatusBadRequest, "file is not 1: "+string(file))
		}
	}

	for idx, file := range req.Files {
		if expected := fmt.Sprintf("%d.txt", idx); file.Filename != expected {
			return nil, apierror.NewAPIError(http.StatusBadRequest, "filename("+expected+") is not correct")
		}

		file, err := e2eutil.ReadMultipartFile(file)

		if err != nil {
			return nil, apierror.NewAPIError(http.StatusBadRequest, err)
		}

		if expected := strconv.Itoa(idx); string(file) != expected {
			return nil, apierror.NewAPIError(http.StatusBadRequest, fmt.Sprintf("file is not %s: %s", expected, string(file)))
		}
	}

	if req.Payload != "payload" {
		return nil, apierror.NewAPIError(http.StatusBadRequest, "payload is not payload")
	}

	if req.Param != "param" {
		return nil, apierror.NewAPIError(http.StatusBadRequest, "param is not param")
	}

	return &types.PostBResponse{
		Message: "OK",
	}, nil
}

// AutoBind - use echo.Bind
func (ctrl *postBController) AutoBind() bool {
	return true
}
