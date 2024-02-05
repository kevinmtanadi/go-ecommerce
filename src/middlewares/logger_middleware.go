package middlewares

import (
	"fmt"
	"go-backend-template/src/constants"
	"io"
	"time"

	"github.com/labstack/echo/v4"
)

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestStart := time.Now()
			err := next(c)
			requestDone := time.Now()

			method := c.Request().Method
			status := c.Response().Status
			request := c.Request().Body
			requestBody := GetRequestBody(request)
			fmt.Printf("[TIME]\t%s\n[%s]\t%s %d %v\n[REQUEST]\n%s\n",
				time.Now().Format(constants.DATE_FORMAT),
				method, c.Path(), status, requestDone.Sub(requestStart),
				requestBody)

			return err
		}
	}
}

func GetRequestBody(request io.ReadCloser) string {
	body, err := io.ReadAll(request)
	if err != nil {
		return ""
	}

	return string(body)
}
