// Package main is a server for sample
package main

import (
	"github.com/go-generalize/api_gen/server_generator/sample/api_gen/bootstrap"
	"github.com/go-generalize/api_gen/server_generator/sample/api_gen/props"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	bootstrap.Bootstrap(&props.ControllerProps{}, e, nil)

	panic(e.Start(":8080"))
}
