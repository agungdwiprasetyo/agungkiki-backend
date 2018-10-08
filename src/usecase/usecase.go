package usecase

import (
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/repository"
)

// InvitationUsecase abstract interface
type InvitationUsecase interface {
	GetAll(offset, limit int) (int, []model.Invitation)
	GetByEmail(email string) *model.Invitation
	GetByName(name string) (int, []model.Invitation)
	GetCount(isAttend bool) (int, error)
	Save(data *model.Invitation) error
	Remove(emails []string) error
}

// NewInvitationUsecase create new usecase
func NewInvitationUsecase(repo repository.InvitationRepository) InvitationUsecase {
	uc := new(invitationUsecase)
	uc.repo = repo
	return uc
}
