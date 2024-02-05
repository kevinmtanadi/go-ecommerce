package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ValidateUserAuth() echo.MiddlewareFunc {
	if os.Getenv("PRODUCTION") == "true" {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				authorization := c.Request().Header.Get("Authorization")
				if authorization == "" {
					return echo.NewHTTPError(401, "Unauthorized")
				}

				parts := strings.Split(authorization, " ")
				if len(parts) != 2 || parts[0] != "Bearer" {
					return echo.NewHTTPError(401, "Unauthorized")
				}

				token := parts[1]
				claims, err := ParseJWT(token)
				if err != nil {
					return echo.NewHTTPError(401, "Unauthorized")
				}

				c.Set("user_id", claims["user_id"].(string))
				c.Set("role_id", claims["role_id"].(string))
				return next(c)
			}
		}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("user_id", "1")
			c.Set("role_id", "1")
			return next(c)
		}
	}
}

func ParseJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
