package main

import (
	controller "github.com/go-generalize/api_gen/v2/samples/standard/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.NewControllers(nil, e)

	// Start echo server

	panic(e.Start(":8080"))
}
