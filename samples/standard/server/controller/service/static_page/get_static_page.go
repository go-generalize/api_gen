// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: devel
// nolint:dupl
package controller

import (
	types "github.com/go-generalize/api_gen/samples/standard/api/service/static_page"
	props "github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// GetStaticPageController ...
type GetStaticPageController interface {
	GetStaticPage(c echo.Context, req *types.GetStaticPageRequest) (res *types.GetStaticPageResponse, err error)
}

type getStaticPageController struct {
	*props.ControllerProps
}

// NewGetStaticPageController ...
func NewGetStaticPageController(cp *props.ControllerProps) GetStaticPageController {
	return &getStaticPageController{
		ControllerProps: cp,
	}
}

// GetStaticPage - GET static_page
func (ctrl *getStaticPageController) GetStaticPage(
	c echo.Context, req *types.GetStaticPageRequest,
) (res *types.GetStaticPageResponse, err error) {
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
func (ctrl *getStaticPageController) AutoBind() bool {
	return true
}
