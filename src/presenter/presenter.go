package presenter

import (
	"net/http"

	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/usecase"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
)

// InvitationPresenter model
type InvitationPresenter struct {
	invitationUsecase usecase.InvitationUsecase
}

// NewInvitationPresenter create new invitation presenter
func NewInvitationPresenter(invitationUsecase usecase.InvitationUsecase) *InvitationPresenter {
	return &InvitationPresenter{invitationUsecase: invitationUsecase}
}

// Mount http router to presenter
func (p *InvitationPresenter) Mount(router *echo.Group) {
	router.GET("/root", p.initGraphqlRoot)
	router.POST("/save", p.save)
}

// InitGraphqlRoot handler
func (p *InvitationPresenter) initGraphqlRoot(c echo.Context) error {
	query := c.QueryParam("query")

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"get_all_invitation": &graphql.Field{
					Name:    "GetAll",
					Type:    graphql.NewList(model.InvitationType),
					Resolve: p.getAll,
				},
				"get_by_email": &graphql.Field{
					Name: "GetByEmail",
					Type: model.InvitationType,
					Args: graphql.FieldConfigArgument{
						"email": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: p.getByEmail,
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
