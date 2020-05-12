// THIS FILE IS A GENERATED CODE. DO NOT EDIT
package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	router *echo.Group
}

func NewRoutes(router *echo.Group) *Routes {
	r := &Routes{
		router: router,
	}
	router.GET("article", r.GetArticle())

	return r
}

func (r *Routes) GetArticle() echo.HandlerFunc {
	i := NewGetArticleController()
	return func(c echo.Context) error {
		req := new(GetArticleRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.GetArticle(c, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	}
}

type IGetArticleController interface {
	GetArticle(c echo.Context, req *GetArticleRequest) (res *GetArticleResponse, err error)
}
