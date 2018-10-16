package presenter

import (
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"github.com/graphql-go/graphql"
)

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
