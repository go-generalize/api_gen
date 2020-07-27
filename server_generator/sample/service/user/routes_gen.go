// THIS FILE IS A GENERATED CODE. DO NOT EDIT
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 0.3.4
package user

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

	router.POST("update_user_name", r.PostUpdateUserName(ctx))
	router.POST("update_user_password", r.PostUpdateUserPassword(ctx))

	return r
}

func (r *Routes) PostUpdateUserName(ctx context.Context) echo.HandlerFunc {
	i := NewPostUpdateUserNameController()
	return func(c echo.Context) error {
		req := new(PostUpdateUserNameRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.PostUpdateUserName(ctx, c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (r *Routes) PostUpdateUserPassword(ctx context.Context) echo.HandlerFunc {
	i := NewPostUpdateUserPasswordController()
	return func(c echo.Context) error {
		req := new(PostUpdateUserPasswordRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.PostUpdateUserPassword(ctx, c, req)
		if err != nil {
			return err
		}
		if res == nil {
			return nil
		}

		return c.JSON(http.StatusOK, res)
	}
}

type IPostUpdateUserNameController interface {
	PostUpdateUserName(c echo.Context, req *PostUpdateUserNameRequest) (res *PostUpdateUserNameResponse, err error)
}

type IPostUpdateUserPasswordController interface {
	PostUpdateUserPassword(c echo.Context, req *PostUpdateUserPasswordRequest) (res *PostUpdateUserPasswordResponse, err error)
}
