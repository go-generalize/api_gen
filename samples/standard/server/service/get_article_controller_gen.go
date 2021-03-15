// Package service ...
// generated version: devel
package service

import (
	"github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

// GetArticleController ...
type GetArticleController struct {
	*props.ControllerProps
}

// NewGetArticleController ...
func NewGetArticleController(cp *props.ControllerProps) *GetArticleController {
	g := &GetArticleController{
		ControllerProps: cp,
	}
	return g
}

// GetArticle ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param ID query integer false ""
// @Success 200 {object} GetArticleResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /service/article [GET]
func (g *GetArticleController) GetArticle(
	c echo.Context, req *GetArticleRequest,
) (res *GetArticleResponse, err error) {
	// API Error Usage: github.com/go-generalize/api_gen/samples/standard/server/wrapper
	//
	// return nil, wrapper.NewAPIError(http.StatusBadRequest)
	//
	// return nil, wrapper.NewAPIError(http.StatusBadRequest).SetError(err)
	//
	// body := map[string]interface{}{
	// 	"code": http.StatusBadRequest,
	// 	"message": "invalid request parameter.",
	// }
	// return nil, wrapper.NewAPIError(http.StatusBadRequest, body).SetError(err)
	panic("require implements.") // FIXME require implements.
}

// AutoBind - use echo.Bind
func (g *GetArticleController) AutoBind() bool {
	return true
}
