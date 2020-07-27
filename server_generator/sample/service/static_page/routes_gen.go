// THIS FILE IS A GENERATED CODE. DO NOT EDIT
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 0.3.4
package static

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	router *echo.Group
}

func NewRoutes(ctx context.Context, router *echo.Group) *Routes {
	r := &Routes{
		router: router,
	}

	router.GET("static_page", r.GetStaticPage(ctx))

	return r
}

func (r *Routes) GetStaticPage(ctx context.Context) echo.HandlerFunc {
	i := NewGetStaticPageController()
	return func(c echo.Context) error {
		req := new(GetStaticPageRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.GetStaticPage(ctx, c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

type IGetStaticPageController interface {
	GetStaticPage(c echo.Context, req *GetStaticPageRequest) (res *GetStaticPageResponse, err error)
}
