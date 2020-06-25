// THIS FILE IS A GENERATED CODE. DO NOT EDIT
package sample

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

	router.POST("create_table", r.PostCreateTable(ctx))
	router.POST("create_user/:id", r.PostCreateUser(ctx))

	return r
}

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

type IPostCreateTableController interface {
	PostCreateTable(c echo.Context, req *PostCreateTableRequest) (res *PostCreateTableResponse, err error)
}

type IPostCreateUserController interface {
	PostCreateUser(c echo.Context, req *PostCreateUserRequest) (res *PostCreateUserResponse, err error)
}