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
	var params model.AllInvitationParam
	params.Page, _ = strconv.Atoi(c.QueryParam("page"))
	params.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	if isAttend, err := strconv.ParseBool(c.QueryParam("isAttend")); err == nil {
		params.IsAttend = &isAttend
	}
	result := p.invitationUsecase.GetAll(&params)
	if result.Error != nil {
		response := helper.NewHTTPResponse(http.StatusInternalServerError, result.Error.Error())
		return response.SetResponse(c)
	}

	meta := helper.Meta{Page: params.Page, Offset: params.Offset, Limit: params.Limit, Total: result.Count}
	response := helper.NewHTTPResponse(http.StatusOK, "success", result.Data, meta)
	return response.SetResponse(c)
}

// GetEvents rest
func (p *InvitationPresenter) GetEvents(c echo.Context) error {
	res := p.invitationUsecase.GetEvent()
	response := helper.NewHTTPResponse(http.StatusOK, "success", res.Data)
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

func (p *InvitationPresenter) login(c echo.Context) error {
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&payload); err != nil {
		response := helper.NewHTTPResponse(http.StatusBadRequest, err.Error())
		return response.SetResponse(c)
	}
	ucRes := p.invitationUsecase.UserLogin(payload.Username, payload.Password)
	if ucRes.Error != nil {
		response := helper.NewHTTPResponse(http.StatusBadRequest, ucRes.Error.Error())
		return response.SetResponse(c)
	}
	response := helper.NewHTTPResponse(http.StatusOK, "success", ucRes.Data)
	return response.SetResponse(c)
}

func (p *InvitationPresenter) saveUser(c echo.Context) error {
	var payload model.User
	if err := c.Bind(&payload); err != nil {
		response := helper.NewHTTPResponse(http.StatusBadRequest, err.Error())
		return response.SetResponse(c)
	}
	ucRes := p.invitationUsecase.SaveUser(&payload)
	if ucRes.Error != nil {
		response := helper.NewHTTPResponse(http.StatusBadRequest, ucRes.Error.Error())
		return response.SetResponse(c)
	}
	response := helper.NewHTTPResponse(http.StatusOK, "success", map[string]interface{}{"userId": ucRes.Data})
	return response.SetResponse(c)
}

func (p *InvitationPresenter) saveEvent(c echo.Context) error {
	var payload model.Event
	if err := c.Bind(&payload); err != nil {
		response := helper.NewHTTPResponse(http.StatusBadRequest, err.Error())
		return response.SetResponse(c)
	}
	err := p.invitationUsecase.SaveEvent(&payload)
	if err != nil {
		response := helper.NewHTTPResponse(http.StatusBadRequest, err.Error())
		return response.SetResponse(c)
	}
	response := helper.NewHTTPResponse(http.StatusOK, "success")
	return response.SetResponse(c)
}
