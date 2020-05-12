// THIS FILE IS A GENERATED CODE. DO NOT EDIT
package sample

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
	router.POST("create_user", r.PostCreateUser())
	router.POST("create_table", r.PostCreateTable())

	return r
}

func (r *Routes) PostCreateUser() echo.HandlerFunc {
	i := NewPostCreateUserController()
	return func(c echo.Context) error {
		req := new(PostCreateUserRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.PostCreateUser(c, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (r *Routes) PostCreateTable() echo.HandlerFunc {
	i := NewPostCreateTableController()
	return func(c echo.Context) error {
		req := new(PostCreateTableRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid request.",
			})
		}
		res, err := i.PostCreateTable(c, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	}
}

type IPostCreateUserController interface {
	PostCreateUser(c echo.Context, req *PostCreateUserRequest) (res *PostCreateUserResponse, err error)
}

type IPostCreateTableController interface {
	PostCreateTable(c echo.Context, req *PostCreateTableRequest) (res *PostCreateTableResponse, err error)
}
