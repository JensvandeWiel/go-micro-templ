package main

import (
	"github.com/JensvandeWiel/go-micro-templ/handlers"
	"github.com/labstack/echo/v4"
)

func createV1Routes(v1 *echo.Group) {
	indexHandler := handlers.NewIndexHandler()
	v1.GET("/hello", indexHandler.HelloWorldHandle)
}
