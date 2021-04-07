// Package controller is for
// generated version: devel
package controller

import (
	types "github.com/go-generalize/api_gen/samples/standard/api/service"
	props "github.com/go-generalize/api_gen/samples/standard/server/props"
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
