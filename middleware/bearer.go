package middleware

import (
	"net/http"
	"strings"

	"github.com/agungdwiprasetyo/agungkiki-backend/helper"
	tokenModule "github.com/agungdwiprasetyo/agungkiki-backend/token"
	"github.com/labstack/echo"
)

// Bearer jwt token middleware
func Bearer(token tokenModule.Token) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				response := helper.NewHTTPResponse(http.StatusUnauthorized, "Invalid Oauth Token")
				return response.SetResponse(c)
			}

			authValues := strings.Split(authorization, " ")
			authType := strings.ToLower(authValues[0])
			if authType != "bearer" || len(authValues) != 2 {
				response := helper.NewHTTPResponse(http.StatusUnauthorized, "Invalid Oauth Token")
				return response.SetResponse(c)
			}

			tokenString := authValues[1]
			isValid, claims := token.Extract(tokenString)
			if !isValid {
				response := helper.NewHTTPResponse(http.StatusUnauthorized, "Invalid Oauth Token")
				return response.SetResponse(c)
			}

			c.Set("tokenClaim", claims)

			return next(c)
		}
	}
}
