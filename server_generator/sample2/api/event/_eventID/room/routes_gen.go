// THIS FILE IS A GENERATED CODE. DO NOT EDIT
package room

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

	router.GET("room", r.GetRoom(ctx))

	return r
}

func (r *Routes) GetRoom(ctx context.Context) echo.HandlerFunc {
	i := NewGetRoomController()
	return func(c echo.Context) error {
		req := new(GetRoomRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.GetRoom(ctx, c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

type IGetRoomController interface {
	GetRoom(c echo.Context, req *GetRoomRequest) (res *GetRoomResponse, err error)
}
