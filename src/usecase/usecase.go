package usecase

import (
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
)

// InvitationUsecase abstract interface
type InvitationUsecase interface {
	GetAll(params *model.AllInvitationParam) UcResult
	GetByWaNumber(waNumber string) *model.Invitation
	GetByName(name string) (int, []model.Invitation)
	GetCount(isAttend bool) (int, error)
	GetEvent() UcResult
	Save(data *model.Invitation) error
	Remove(numbers []string) error
	UserLogin(username, password string) UcResult
	SaveUser(userData *model.User) UcResult
	SaveEvent(data *model.Event) error
}
