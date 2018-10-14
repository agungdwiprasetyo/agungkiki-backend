package middleware

import (
	"net/http"

	"github.com/agungdwiprasetyo/agungkiki-backend/helper"
	tokenModule "github.com/agungdwiprasetyo/agungkiki-backend/token"
	"github.com/labstack/echo"
)

// Role middleware
func Role(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			sess, _ := c.Get("tokenClaim").(*tokenModule.Claim)
			always := false
			for _, role := range roles {
				always := sess.Audience == role
				if always {
					break
				}
			}
			if sess.Audience == "admin" {
				always = true
			}

			if !always {
				response := helper.NewHTTPResponse(http.StatusForbidden, "You can't access it")
				return response.SetResponse(c)
			}

			return next(c)
		}
	}
}
