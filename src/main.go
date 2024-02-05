package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"

	docs "go-backend-template/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	app := echo.New()

	module := Module{}
	module.New(app)

	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8000"
	}

	if os.Getenv("PRODUCTION") == "false" {
		docs.SwaggerInfo.Host = os.Getenv("APP_HOST")
		app.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", port)))
}
