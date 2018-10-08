package presenter

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/agungdwiprasetyo/agungkiki-backend/helper"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"github.com/agungdwiprasetyo/go-utils/debug"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/labstack/echo"
)

// Save rest
func (p *InvitationPresenter) save(c echo.Context) error {
	response := new(helper.HTTPResponse)
	response.Success = true
	response.Code = http.StatusOK
	response.Message = "Success save"

	var errs error
	errs = multierror.Append(errs, fmt.Errorf("mantul"))

	var payload model.Invitation
	if err := c.Bind(&payload); err != nil {
		response.Success = false
		response.Code = http.StatusBadRequest
		response.Message = "error"
		response.Errors = append(response.Errors, err.Error())
		return response.SetResponse(c)
	}

	if err := p.invitationUsecase.Save(&payload); err != nil {
		response.Success = false
		response.Code = http.StatusBadRequest
		response.Message = "error"
		response.Errors = append(response.Errors, err.Error())
		if strings.Contains(err.Error(), "duplicate") {
			response.Code = http.StatusConflict
			response.Message = fmt.Sprintf("Email %s telah digunakan", payload.Email)
		}
		errs = multierror.Append(errs, err)
		debug.Println(errs)
		return response.SetResponse(c)
	}

	return response.SetResponse(c)
}
