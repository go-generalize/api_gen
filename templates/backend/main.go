package main

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-generalize/api_gen/templates/backend/interfaces"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ctx := context.Background()

	e := echo.New()

	middlewareList := make([]*interfaces.MiddlewareSet, 0)
	mid := &interfaces.MiddlewareSet{
		Path: "/api/",
		MiddlewareFunc: []echo.MiddlewareFunc{
			middleware.JWTWithConfig(middleware.JWTConfig{
				ContextKey:     "jwt",
				SuccessHandler: nil,
				SigningKey:     []byte("key"),
				SigningMethod:  jwt.SigningMethodHS512.Name,
				Claims:         new(jwt.StandardClaims),
				TokenLookup:    "cookie:ApiGenSession",
			}),
		},
	}
	middlewareList = append(middlewareList, mid)

	e.Debug = true
	e.Use(middleware.Recover())

	interfaces.Bootstrap(ctx, e, nil)

	fmt.Println("All routes are...")
	for _, r := range e.Routes() {
		fmt.Printf("%s %s: %s\n", r.Method, r.Path, r.Name)
	}

	if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}
