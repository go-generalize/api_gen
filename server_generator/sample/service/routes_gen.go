// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 0.3.5

package service

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Routes ...
type Routes struct {
	router *echo.Group
}

// NewRoutes ...
func NewRoutes(ctx context.Context, router *echo.Group) *Routes {
	r := &Routes{
		router: router,
	}

	router.GET("article", r.GetArticle(ctx))

	return r
}

// GetArticle ...
func (r *Routes) GetArticle(ctx context.Context) echo.HandlerFunc {
	i := NewGetArticleController()
	return func(c echo.Context) error {
		req := new(GetArticleRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.GetArticle(ctx, c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// IGetArticleController ...
type IGetArticleController interface {
	GetArticle(c echo.Context, req *GetArticleRequest) (res *GetArticleResponse, err error)
}
