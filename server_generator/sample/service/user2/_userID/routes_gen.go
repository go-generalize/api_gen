// Code generated by server_generator. DO NOT EDIT.
// generated version: devel

package _userID

import (
	"io"
	"log"
	"net/http"

	"github.com/go-generalize/api_gen/server_generator/sample/props"
	"github.com/go-generalize/api_gen/server_generator/sample/wrapper"
	"github.com/labstack/echo/v4"
)

// Routes ...
type Routes struct {
	router *echo.Group
}

// NewRoutes ...
func NewRoutes(p *props.ControllerProps, router *echo.Group, opts ...io.Writer) *Routes {
	if len(opts) > 0 {
		if w := opts[0]; w != nil {
			log.SetOutput(w)
		}
	}
	r := &Routes{
		router: router,
	}
	router.GET("user_job_get", r.GetUserJobGet(p))
	return r
}

// GetUserJobGet ...
func (r *Routes) GetUserJobGet(p *props.ControllerProps) echo.HandlerFunc {
	i := NewGetUserJobGetController(p)
	return func(c echo.Context) error {
		req := new(GetUserJobGetRequest)
		if err := c.Bind(req); err != nil {
			log.Printf("failed to JSON binding(/service/user2/{userID}/user_job_get): %+v", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.GetUserJobGet(c, req)
		if err != nil {
			if werr, ok := err.(*wrapper.APIError); ok {
				log.Printf("%+v", werr)
				return c.JSON(werr.Status, werr.Body)
			}
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// IGetUserJobGetController ...
type IGetUserJobGetController interface {
	GetUserJobGet(c echo.Context, req *GetUserJobGetRequest) (res *GetUserJobGetResponse, err error)
}
