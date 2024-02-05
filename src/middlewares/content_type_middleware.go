package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ValidateContentType() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			contentType := c.Request().Header.Get("Content-Type")
			if c.Request().Method != http.MethodGet && contentType != "application/json" {
				return echo.NewHTTPError(http.StatusBadRequest, "The Content-Type header is not valid. Please use application/json")
			}
			return next(c)
		}
	}
}
