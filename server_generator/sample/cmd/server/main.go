// Package main is a server for sample
package main

import (
	"github.com/go-generalize/api_gen/server_generator/sample/apigen"
	"github.com/go-generalize/api_gen/server_generator/sample/apigen/props"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	apigen.Bootstrap(&props.ControllerProps{}, e, nil)

	panic(e.Start(":8080"))
}
