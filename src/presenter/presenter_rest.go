package presenter

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/agungdwiprasetyo/agungkiki-backend/helper"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"github.com/agungdwiprasetyo/go-utils/debug"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/labstack/echo"
)

// GetAll rest
func (p *InvitationPresenter) GetAll(c echo.Context) error {
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	count, data := p.invitationUsecase.GetAll(offset, limit)

	meta := helper.Meta{Offset: offset, Limit: limit, Total: count}
	response := helper.NewHTTPResponse(http.StatusOK, "success", data, meta)
	return response.SetResponse(c)
}

// Save rest
func (p *InvitationPresenter) Save(c echo.Context) error {
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
		return response.SetResponse(c)
	}

	if err := p.invitationUsecase.Save(&payload); err != nil {
		response.Success = false
		response.Code = http.StatusBadRequest
		response.Message = "error"
		if strings.Contains(err.Error(), "duplicate") {
			response.Code = http.StatusConflict
			response.Message = fmt.Sprintf("Nomor %s telah mengisi data", payload.WaNumber)
		}
		errs = multierror.Append(errs, err)
		debug.Println(errs)
		return response.SetResponse(c)
	}

	return response.SetResponse(c)
}

// Remove rest
func (p *InvitationPresenter) Remove(c echo.Context) error {
	response := new(helper.HTTPResponse)
	response.Success = true
	response.Code = http.StatusOK
	response.Message = "Success remove"

	var payload []string
	if err := c.Bind(&payload); err != nil {
		response.Success = false
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
		return response.SetResponse(c)
	}

	if err := p.invitationUsecase.Remove(payload); err != nil {
		response.Success = false
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
		return response.SetResponse(c)
	}

	return response.SetResponse(c)
}
