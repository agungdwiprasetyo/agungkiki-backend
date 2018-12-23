package helper

import (
	"net/http"

	"github.com/labstack/echo"
)

// HTTPResponse model
type HTTPResponse struct {
	Code    int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Meta    *Meta       `json:"meta,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Meta model
type Meta struct {
	Page   int `json:"page"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"totalRecords"`
}

// NewHTTPResponse constructor
func NewHTTPResponse(code int, message string, params ...interface{}) *HTTPResponse {
	commonResponse := new(HTTPResponse)

	for i, param := range params {
		if i == 0 {
			commonResponse.Data = param
		} else if i == 1 {
			if meta, ok := param.(Meta); ok {
				commonResponse.Meta = &meta
			}
		}
	}

	if code < http.StatusBadRequest {
		commonResponse.Success = true
	}

	commonResponse.Code = code
	commonResponse.Message = message
	return commonResponse
}

// SetResponse method
func (h *HTTPResponse) SetResponse(c echo.Context) error {
	if !h.Success {
		h.Error = h.Data
		h.Data = nil
	}
	return c.JSON(h.Code, h)
}
