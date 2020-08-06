// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 0.3.5

package sample

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

	router.POST("create_table", r.PostCreateTable(ctx))
	router.POST("create_user", r.PostCreateUser(ctx))

	return r
}

// PostCreateTable ...
func (r *Routes) PostCreateTable(ctx context.Context) echo.HandlerFunc {
	i := NewPostCreateTableController()
	return func(c echo.Context) error {
		req := new(PostCreateTableRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.PostCreateTable(ctx, c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// PostCreateUser ...
func (r *Routes) PostCreateUser(ctx context.Context) echo.HandlerFunc {
	i := NewPostCreateUserController()
	return func(c echo.Context) error {
		req := new(PostCreateUserRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.PostCreateUser(ctx, c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

// IPostCreateTableController ...
type IPostCreateTableController interface {
	PostCreateTable(c echo.Context, req *PostCreateTableRequest) (res *PostCreateTableResponse, err error)
}

// IPostCreateUserController ...
type IPostCreateUserController interface {
	PostCreateUser(c echo.Context, req *PostCreateUserRequest) (res *PostCreateUserResponse, err error)
}
