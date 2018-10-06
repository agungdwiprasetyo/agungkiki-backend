package middleware

import (
	"net/http"

	"github.com/labstack/echo"
)

// SetCORS middleware
func SetCORS() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")

			if c.Request().Method == http.MethodOptions {
				return c.JSON(200, nil)
			}
			return next(c)
		}
	}
}
