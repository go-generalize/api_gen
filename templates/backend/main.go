package main

import (
	"fmt"

	"github.com/go-generalize/api_gen/templates/backend/props"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Debug = true
	e.Use(echoMiddleware.Recover())

	ctrl := NewControllers(new(props.ControllerProps), e)

	ctrl.AddMiddleware("/api/", echoMiddleware.JWTWithConfig(echoMiddleware.JWTConfig{
		ContextKey:     "jwt",
		SuccessHandler: nil,
		SigningKey:     []byte("key"),
		SigningMethod:  jwt.SigningMethodHS512.Name,
		Claims:         new(jwt.StandardClaims),
		TokenLookup:    "cookie:ApiGenSession",
	}))

	fmt.Println("All routes are...")
	for _, r := range e.Routes() {
		fmt.Printf("%s %s: %s\n", r.Method, r.Path, r.Name)
	}

	if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}
