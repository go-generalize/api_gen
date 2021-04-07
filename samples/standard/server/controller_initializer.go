// Package controller binds handlers to echo
// generated version: devel
package controller

import (
	"net/http"
	"runtime/debug"
	"strings"

	props "github.com/go-generalize/api_gen/samples/standard/server/props"
	"github.com/labstack/echo/v4"
)

type middleware struct {
	path       string
	middleware echo.MiddlewareFunc
}

// Controllers binds handlers to echo
type Controllers struct {
	middlewares []middleware
}

// NewControllers returns a new Controllers
func NewControllers(
	props *props.ControllerProps, e *echo.Echo,
) *Controllers {
	ctrl := &Controllers{}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, m := range ctrl.middlewares {
				if strings.HasPrefix(c.Request().URL.Path, m.path) {
					next = m.middleware(next)
				}
			}

			return next(c)
		}
	})
	e.Use(func(before echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() {
				recoverErr := recover()
				if recoverErr == nil {
					return
				}

				debug.PrintStack()

				if httpErr, ok := recoverErr.(*echo.HTTPError); ok {
					err = c.JSON(httpErr.Code, httpErr.Message)
				}

				err = c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "internal server error.",
				})
			}()

			return before(c)
		}
	})

	addRoutes(e, props)

	return ctrl
}

// AddMiddleware adds 'm' to the paths starting with the 'path'.
func (c *Controllers) AddMiddleware(path string, m echo.MiddlewareFunc) {
	c.middlewares = append(c.middlewares, middleware{
		path:       path,
		middleware: m,
	})
}
