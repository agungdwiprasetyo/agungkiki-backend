package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/agungdwiprasetyo/agungkiki-backend/helper"
	"github.com/labstack/echo"
)

var (
	green        = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white        = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow       = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	red          = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue         = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta      = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan         = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset        = string([]byte{27, 91, 48, 109})
	disableColor = false
)

// Logger middleware
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()
			res := c.Response()

			err := next(c)
			if err != nil {
				response := helper.NewHTTPResponse(http.StatusNotFound, fmt.Sprintf("Route '%s %s' Not Found", req.Method, req.URL.String()))
				err = response.SetResponse(c)
			}
			end := time.Now()

			statusColor := colorForStatus(res.Status)
			methodColor := colorForMethod(req.Method)
			resetColor := reset

			fmt.Fprintf(os.Stdout, "%s[LOGGER]%s %v | %s %3d %s | %13v | %15s | %s %-7s %s %s\n",
				white, resetColor,
				end.Format("2006/01/02 - 15:04:05"),
				statusColor, res.Status, resetColor,
				end.Sub(start),
				c.RealIP(),
				methodColor, req.Method, resetColor,
				req.RequestURI,
			)
			return err
		}
	}
}

func colorForStatus(code int) string {
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}
