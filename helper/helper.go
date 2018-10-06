package helper

import "github.com/labstack/echo"

// HTTPResponse model
type HTTPResponse struct {
	Code    int         `json:"-"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

// NewHTTPResponse constructor
func NewHTTPResponse(code int, success bool, message string, data interface{}, errors []string) *HTTPResponse {
	return &HTTPResponse{
		Code:    code,
		Success: success,
		Message: message,
		Data:    data,
		Errors:  errors,
	}
}

// SetResponse method
func (h *HTTPResponse) SetResponse(c echo.Context) error {
	return c.JSON(h.Code, h)
}
