package sample

import (
	"context"
	"log"
	"net/http"

	p1 "github.com/go-generalize/api_gen/server_generator/sample/service"
	p2 "github.com/go-generalize/api_gen/server_generator/sample/service/user"
	p3 "github.com/go-generalize/api_gen/server_generator/sample/service/user2"
	"github.com/labstack/echo/v4"
)

type MiddlewareList []*MiddlewareSet
type MiddlewareMap map[string][]echo.MiddlewareFunc

type MiddlewareSet struct {
	Path           string
	MiddlewareFunc []echo.MiddlewareFunc
}

func (m MiddlewareList) ToMap() MiddlewareMap {
	mf := make(map[string][]echo.MiddlewareFunc)
	for _, middleware := range m {
		mf[middleware.Path] = middleware.MiddlewareFunc
	}
	return mf
}

func Bootstrap(ctx context.Context, e *echo.Echo, middlewareList MiddlewareList) {
	middleware := middlewareList.ToMap()

	// error handling
	e.Use(func(before echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() {
				recoverErr := recover()
				if recoverErr == nil {
					return
				}
				if httpErr, ok := recoverErr.(*echo.HTTPError); ok {
					err = c.JSON(httpErr.Code, httpErr.Message)
				}
				log.Printf("panic: %#v", recoverErr)
				err = c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "internal server error.",
				})
			}()

			err = before(c)
			return err
		}
	})

	g0 := e.Group("")
	setMiddleware(g0, "/", middleware)
	NewRoutes(ctx, g0)

	g1 := g0.Group("service/")
	setMiddleware(g1, "/service/", middleware)
	p1.NewRoutes(ctx, g1)

	g2 := g1.Group("user/")
	setMiddleware(g2, "/service/user/", middleware)
	p2.NewRoutes(ctx, g2)

	g3 := g1.Group("user2/")
	setMiddleware(g3, "/service/user2/", middleware)
	p3.NewRoutes(ctx, g3)
}

func setMiddleware(group *echo.Group, path string, list MiddlewareMap) {
	if ms, ok := list[path]; ok {
		for _, m := range ms {
			group.Use(m)
		}
	}
}
