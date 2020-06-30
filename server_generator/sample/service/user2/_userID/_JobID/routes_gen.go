// THIS FILE IS A GENERATED CODE. DO NOT EDIT
package _JobID

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

	router.PUT("job", r.PutJob(ctx))

	return r
}

func (r *Routes) PutJob(ctx context.Context) echo.HandlerFunc {
	i := NewPutJobController()
	return func(c echo.Context) error {
		req := new(PutJobRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.PutJob(ctx, c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

type IPutJobController interface {
	PutJob(c echo.Context, req *PutJobRequest) (res *PutJobResponse, err error)
}
