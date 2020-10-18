// +build mock
// Code generated by server_generator. DO NOT EDIT.
// generated version: unknown

package sample2

import (
	"log"
	"net/http"
	"path/filepath"

	apiEventEventIDRoom "github.com/go-generalize/api_gen/server_generator/sample2/api/event/_eventID/room"
	"github.com/labstack/echo/v4"

	props "github.com/go-generalize/api_gen/server_generator/sample2/props"
)

// Bootstrap ...
func MockBootstrap(p *props.ControllerProps, e *echo.Echo, jsonDir string) {
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

	apiEventEventIDRoomGroup := rootGroup.Group("api/event/:eventID/room/")
	apiEventEventIDRoom.NewMockRoutes(p, apiEventEventIDRoomGroup, filepath.Join(jsonDir, "/api/event/_eventID/room"))
}