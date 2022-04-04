// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: devel
// nolint:dupl
package controller

import (
	props "github.com/go-generalize/api_gen/v2/pkg/server/tmp/output/props"
	types "github.com/go-generalize/api_gen/v2/samples/standard/api/service"
	"github.com/labstack/echo/v4"
)

// GetArticleController ...
type GetArticleController interface {
	GetArticle(c echo.Context, req *types.GetArticleRequest) (res *types.GetArticleResponse, err error)
}

type getArticleController struct {
	*props.ControllerProps
}

// NewGetArticleController ...
func NewGetArticleController(cp *props.ControllerProps) GetArticleController {
	return &getArticleController{
		ControllerProps: cp,
	}
}

// GetArticle - GET article
func (ctrl *getArticleController) GetArticle(
	c echo.Context, req *types.GetArticleRequest,
) (res *types.GetArticleResponse, err error) {
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
func (ctrl *getArticleController) AutoBind() bool {
	return true
}
