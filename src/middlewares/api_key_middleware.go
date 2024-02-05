package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
)

func ValidateApiKey() echo.MiddlewareFunc {
	if os.Getenv("API_KEY") != "" {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				if c.Path() == "/swagger/*" {
					return next(c)
				}

				apiKey := c.Request().Header.Get("X-ACCESS-TOKEN")
				if apiKey != os.Getenv("API_KEY") {
					return echo.NewHTTPError(401, "API Key invalid")
				} else {
					return next(c)
				}
			}
		}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
