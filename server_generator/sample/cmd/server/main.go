package main

import (
	"github.com/go-generalize/api_gen/server_generator/sample"
	"github.com/go-generalize/api_gen/server_generator/sample/props"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	sample.Bootstrap(&props.ControllerProps{}, e, nil)

	e.Start(":8080")
}
