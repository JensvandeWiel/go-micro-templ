package main

import (
	"flag"
	"fmt"
	_ "github.com/JensvandeWiel/go-micro-templ/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log/slog"
)

// @title go-micro-templ
// @version 1.0
// @description Go microservice template
// termsOfService http://swagger.io/terms/

// @contact.name Jens van de Wiel
// @contact.url https://jens.vandewiel.eu
// @contact.email jens.vdwiel@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /v1
func main() {
	logLevel := flag.String("log-level", "info", "Log level")
	configPath := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	slog.SetLogLoggerLevel(convertLogLevel(*logLevel))

	// Load the config
	config, err := GetConfig(*configPath)
	if err != nil {
		panic(err)
	}

	if config.Environment == "development" {
		slog.Debug("Running in development mode", slog.String("config", fmt.Sprintf("%+v", config)))
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(slogecho.New(slog.Default()))
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")
	createV1Routes(v1)

	if config.Environment != "production" {
		slog.Debug("Environment is not production, enabling swagger endpoint")
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
	slog.Info("Starting server", slog.String("port", config.Port), slog.String("host", config.Host))
	e.Logger.Fatal(e.Start(config.Host + ":" + config.Port))
}
