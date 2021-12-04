package main

import (
	"log"

	controller "github.com/go-generalize/api_gen/v2/samples/empty_root/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	ctrl := controller.NewControllers(nil, e)

	ctrl.DisableErrorHandling()
	ctrl.AddMiddleware("/", echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				log.Println("logging", err)
				return c.JSON(500, err)
			}

			return nil
		}
	}))

	// Start echo server
	panic(e.Start(":5050"))
}
