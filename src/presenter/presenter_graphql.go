package presenter

import (
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"github.com/graphql-go/graphql"
)

func (p *InvitationPresenter) initGraphQlSchema() (graphql.Schema, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"get_all_invitation": &graphql.Field{
					Name: "GetAll",
					Type: graphql.NewList(new(model.Invitation).MakeObject()),
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
					Type: new(model.Invitation).MakeObject(),
					Args: graphql.FieldConfigArgument{
						"wa_number": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: p.getByWaNumber,
				},
				"get_by_name": &graphql.Field{
					Name: "GetByName",
					Type: graphql.NewList(new(model.Invitation).MakeObject()),
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
				"get_event": &graphql.Field{
					Name:    "GetEvent",
					Type:    new(model.Event).MakeObject(),
					Resolve: p.getEvent,
				},
				"get_visitor": &graphql.Field{
					Name: "GetVisitor",
					Type: graphql.NewList(new(model.Visitor).MakeObject()),
					Args: graphql.FieldConfigArgument{
						"from_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"to_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: p.getVisitor,
				},
			},
		}),
	})
	return schema, err
}

// GetAll graphql query
func (p *InvitationPresenter) getAll(params graphql.ResolveParams) (interface{}, error) {
	var inParams model.AllInvitationParam
	inParams.Offset, _ = params.Args["offset"].(int)
	inParams.Limit, _ = params.Args["limit"].(int)
	if isAttend, ok := params.Args["is_attend"].(bool); ok {
		inParams.IsAttend = &isAttend
	}
	res := p.invitationUsecase.GetAll(&inParams)
	if res.Error != nil {
		return nil, res.Error
	}
	return res.Data, nil
}

// GetByWaNumber graphql query
func (p *InvitationPresenter) getByWaNumber(params graphql.ResolveParams) (interface{}, error) {
	waNumber, _ := params.Args["wa_number"].(string)
	data := p.invitationUsecase.GetByWaNumber(waNumber)
	return data, nil
}

// GetByName graphql query
func (p *InvitationPresenter) getByName(params graphql.ResolveParams) (interface{}, error) {
	name, _ := params.Args["name"].(string)
	_, data := p.invitationUsecase.GetByName(name)
	return data, nil
}

// getTotalPresent graphql query
func (p *InvitationPresenter) getCount(params graphql.ResolveParams) (interface{}, error) {
	isAttend, _ := params.Args["is_attend"].(bool)
	return p.invitationUsecase.GetCount(isAttend)
}

func (p *InvitationPresenter) getEvent(params graphql.ResolveParams) (interface{}, error) {
	result := p.invitationUsecase.GetEvent()
	return result.Data, result.Error
}

func (p *InvitationPresenter) getVisitor(params graphql.ResolveParams) (interface{}, error) {
	fromDate, _ := params.Args["from_date"].(string)
	toDate, _ := params.Args["to_date"].(string)
	result := p.invitationUsecase.GetVisitor(fromDate, toDate)
	return result.Data, result.Error
}
