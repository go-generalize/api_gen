// THIS FILE IS A GENERATED CODE. DO NOT EDIT
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 0.3.4
package _userID

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

	router.GET("user_job_get", r.GetUserJobGet(ctx))

	return r
}

func (r *Routes) GetUserJobGet(ctx context.Context) echo.HandlerFunc {
	i := NewGetUserJobGetController()
	return func(c echo.Context) error {
		req := new(GetUserJobGetRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.GetUserJobGet(ctx, c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

type IGetUserJobGetController interface {
	GetUserJobGet(c echo.Context, req *GetUserJobGetRequest) (res *GetUserJobGetResponse, err error)
}
