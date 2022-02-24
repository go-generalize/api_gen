package main

import (
	controller "github.com/go-generalize/api_gen/v2/samples/standard/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	controller.NewControllers(nil, e)

	// Start echo server

	panic(e.Start(":8080"))
}
