package sample

import (
	"context"
	"log"
	"net/http"

	"github.com/go-generalize/api_gen/server_generator/sample/service"
	serviceUser "github.com/go-generalize/api_gen/server_generator/sample/service/user"
	serviceUser2 "github.com/go-generalize/api_gen/server_generator/sample/service/user2"
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

	rootGroup := e.Group("")
	setMiddleware(rootGroup, "/", middleware)
	NewRoutes(ctx, rootGroup)

	serviceGroup := rootGroup.Group("service/")
	setMiddleware(serviceGroup, "/service/", middleware)
	service.NewRoutes(ctx, serviceGroup)

	serviceUserGroup := serviceGroup.Group("user/")
	setMiddleware(serviceUserGroup, "/service/user/", middleware)
	serviceUser.NewRoutes(ctx, serviceUserGroup)

	serviceUser2Group := serviceGroup.Group("user2/")
	setMiddleware(serviceUser2Group, "/service/user2/", middleware)
	serviceUser2.NewRoutes(ctx, serviceUser2Group)
}

func setMiddleware(group *echo.Group, path string, list MiddlewareMap) {
	if ms, ok := list[path]; ok {
		for _, m := range ms {
			group.Use(m)
		}
	}
}
