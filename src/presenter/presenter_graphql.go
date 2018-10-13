package presenter

import (
	"github.com/graphql-go/graphql"
)

// GetAll graphql query
func (p *InvitationPresenter) getAll(params graphql.ResolveParams) (interface{}, error) {
	offset, _ := params.Args["offset"].(int)
	limit, _ := params.Args["limit"].(int)
	_, data := p.invitationUsecase.GetAll(offset, limit)
	return data, nil
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
