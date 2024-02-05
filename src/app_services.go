package main

import (
	"go-backend-template/src/middlewares"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Service struct {
	app *echo.Echo
}

func NewService(app *echo.Echo) *Service {
	return &Service{
		app: app,
	}
}

func (s *Service) UseMiddlewares() {
	godotenv.Load(".env")

	s.app.Use(middleware.Recover())
	s.app.Use(middlewares.Logger())
	s.app.Use(middlewares.ValidateContentType())
	s.app.Use(middleware.Secure())
	s.app.Use(middlewares.ValidateApiKey())
}
