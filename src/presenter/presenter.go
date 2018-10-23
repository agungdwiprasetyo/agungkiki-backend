package presenter

import (
	"net/http"

	"github.com/agungdwiprasetyo/agungkiki-backend/middleware"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/usecase"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
)

// InvitationPresenter model
type InvitationPresenter struct {
	invitationUsecase usecase.InvitationUsecase
	bearerMiddleware  echo.MiddlewareFunc
}

// NewInvitationPresenter create new invitation presenter
func NewInvitationPresenter(invitationUsecase usecase.InvitationUsecase, mid echo.MiddlewareFunc) *InvitationPresenter {
	return &InvitationPresenter{invitationUsecase: invitationUsecase, bearerMiddleware: mid}
}

// Mount http router to presenter
func (p *InvitationPresenter) Mount(router *echo.Group) {
	router.GET("/root", p.initGraphqlRoot)

	router.GET("/all", p.GetAll, p.bearerMiddleware)
	router.GET("/event", p.GetEvents)
	router.POST("/save", p.Save)
	router.DELETE("/remove", p.Remove, p.bearerMiddleware, middleware.Role())

	router.POST("/user/login", p.login)
	router.POST("/user/new", p.saveUser, p.bearerMiddleware)
}

// InitGraphqlRoot handler
func (p *InvitationPresenter) initGraphqlRoot(c echo.Context) error {
	query := c.QueryParam("query")

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"get_all_invitation": &graphql.Field{
					Name: "GetAll",
					Type: graphql.NewList(model.InvitationType),
					Args: graphql.FieldConfigArgument{
						"offset": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: p.getAll,
				},
				"get_by_wa_number": &graphql.Field{
					Name: "GetByWaNumber",
					Type: model.InvitationType,
					Args: graphql.FieldConfigArgument{
						"wa_number": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: p.getByWaNumber,
				},
				"get_by_name": &graphql.Field{
					Name: "GetByName",
					Type: graphql.NewList(model.InvitationType),
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: p.getByName,
				},
				"get_count": &graphql.Field{
					Name: "GetCount",
					Type: graphql.Int,
					Args: graphql.FieldConfigArgument{
						"is_attend": &graphql.ArgumentConfig{
							Type: graphql.Boolean,
						},
					},
					Resolve: p.getCount,
				},
			},
		}),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if result.HasErrors() {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": result.Errors})
	}

	return c.JSON(http.StatusOK, result)
}
