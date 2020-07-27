// THIS FILE IS A GENERATED CODE. DO NOT EDIT
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 0.3.4

package sample2

import (
	"context"
	"log"
	"net/http"

	apiEventEventIDRoom "github.com/go-generalize/api_gen/server_generator/sample2/api/event/_eventID/room"
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

	rootGroup := e.Group("/")
	setMiddleware(rootGroup, "/", middleware)

	apiEventEventIDRoomGroup := rootGroup.Group("api/event/:eventID/room/")
	setMiddleware(apiEventEventIDRoomGroup, "/api/event/:eventID/room/", middleware)
	apiEventEventIDRoom.NewRoutes(ctx, apiEventEventIDRoomGroup)
}

func setMiddleware(group *echo.Group, path string, list MiddlewareMap) {
	if ms, ok := list[path]; ok {
		for _, m := range ms {
			group.Use(m)
		}
	}
}
