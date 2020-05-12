package service

import (
	"github.com/labstack/echo/v4"
)

type GetArticleController struct {
}

func NewGetArticleController() *GetArticleController {
	g := &GetArticleController{}
	return g
}

func (g *GetArticleController) GetArticle(c echo.Context,
	req *GetArticleRequest) (
	res *GetArticleResponse,
	err error) {
	return &GetArticleResponse{
		ID:    req.ID * 2,
		Group: []string{"a", "b", "c"},
		Body:  "hogehoge",
	}, nil
}
