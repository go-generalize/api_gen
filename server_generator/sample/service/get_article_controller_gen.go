// Package service ...
// generated version: devel
package service

import (
	"github.com/go-generalize/api_gen/server_generator/sample/props"
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
// @Param ID query integer WIP:${isRequire} WIP:${description}
// @Success 200 {object} GetArticleResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /service/article [GET]
func (g *GetArticleController) GetArticle(
	_ echo.Context, req *GetArticleRequest,
) (res *GetArticleResponse, err error) {
	return &GetArticleResponse{
		ID:    req.ID * 2,
		Group: []string{"a", "b", "c"},
		Body:  "hogehoge",
	}, nil
}
