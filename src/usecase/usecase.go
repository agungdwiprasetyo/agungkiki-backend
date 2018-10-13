package usecase

import (
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/repository"
	tokenModule "github.com/agungdwiprasetyo/agungkiki-backend/token"
)

// InvitationUsecase abstract interface
type InvitationUsecase interface {
	GetAll(offset, limit int) (int, []model.Invitation)
	GetByWaNumber(waNumber string) *model.Invitation
	GetByName(name string) (int, []model.Invitation)
	GetCount(isAttend bool) (int, error)
	Save(data *model.Invitation) error
	Remove(numbers []string) error
}

// NewInvitationUsecase create new usecase
func NewInvitationUsecase(token *tokenModule.Token, repo repository.InvitationRepository) InvitationUsecase {
	uc := new(invitationUsecase)
	uc.repo = repo
	return uc
}
