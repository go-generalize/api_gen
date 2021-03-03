// +build mock
// Code generated by server_generator. DO NOT EDIT.
// generated version: devel

// Package main is a mock server
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-generalize/api_gen/server_generator/sample2/apigen"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := flag.Int("port", 15000, "mock server listen port")
	jsonDir := flag.String("json-dir", "server_generator/sample2/apigen/mock_jsons", "mock jsons directory")
	flag.Parse()

	jsonDirPath := *jsonDir
	if !strings.HasSuffix(jsonDirPath, "/") {
		jsonDirPath += "/"
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		serverName := fmt.Sprintf("api_gen-mock_server/%s", "devel")
		return func(c echo.Context) error {
			c.Response().Header().Set("Server", serverName)
			err := next(c)
			return err
		}
	})

	apigen.MockBootstrap(e, os.Stderr, jsonDirPath)

	fmt.Println("All routes are...")
	for _, r := range e.Routes() {
		fmt.Printf("%s %s: %s\n", r.Method, r.Path, r.Name)
	}

	l := fmt.Sprintf(":%d", *port)
	log.Fatal(e.Start(l))
}
