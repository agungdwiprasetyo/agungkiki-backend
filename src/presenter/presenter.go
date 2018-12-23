package presenter

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/agungdwiprasetyo/agungkiki-backend/helper"
	"github.com/agungdwiprasetyo/agungkiki-backend/middleware"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/usecase"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
)

// InvitationPresenter model
type InvitationPresenter struct {
	invitationUsecase usecase.InvitationUsecase
	bearerMiddleware  echo.MiddlewareFunc
	graphqlSchema     graphql.Schema
}

// NewInvitationPresenter create new invitation presenter
func NewInvitationPresenter(invitationUsecase usecase.InvitationUsecase, mid echo.MiddlewareFunc) *InvitationPresenter {
	return &InvitationPresenter{invitationUsecase: invitationUsecase, bearerMiddleware: mid}
}

// Mount http router to presenter
func (p *InvitationPresenter) Mount(router *echo.Group) {
	p.graphqlSchema, _ = p.initGraphQlSchema()

	router.Any("/graphql", p.handleGraphqlRoot)

	router.GET("/auth", p.auth, p.bearerMiddleware)
	router.GET("/all", p.GetAll, p.bearerMiddleware)
	router.GET("/event", p.GetEvents)
	router.POST("/event/save", p.saveEvent, p.bearerMiddleware)
	router.POST("/save", p.Save)
	router.DELETE("/remove", p.Remove, p.bearerMiddleware, middleware.Role())

	router.POST("/user/login", p.login)
	router.POST("/user/new", p.saveUser, p.bearerMiddleware)
}

func (p *InvitationPresenter) auth(c echo.Context) error {
	response := helper.NewHTTPResponse(http.StatusOK, "ok")
	return response.SetResponse(c)
}

// InitGraphqlRoot handler
func (p *InvitationPresenter) handleGraphqlRoot(c echo.Context) error {
	var query string
	if query = c.QueryParam("query"); query == "" {
		queryBody, _ := ioutil.ReadAll(c.Request().Body)
		query = string(queryBody)
	}

	result := graphql.Do(graphql.Params{
		Schema:        p.graphqlSchema,
		RequestString: query,
	})
	if result.HasErrors() {
		response := helper.NewHTTPResponse(http.StatusBadRequest, "Query error", result.Errors)
		return response.SetResponse(c)
	}

	if strings.Contains(query, "get_count") {
		p.invitationUsecase.AddVisitor(c.RealIP(), query)
	}

	return c.JSON(http.StatusOK, result)
}
