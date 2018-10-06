package usecase

import (
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/repository"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

type invitationUsecase struct {
	repo repository.InvitationRepository
}

func (uc *invitationUsecase) GetAll() (int, []model.Invitation) {
	result := <-uc.repo.FindAll()
	if result.Error != nil {
		return 0, nil
	}
	data, _ := result.Data.([]model.Invitation)
	return result.Count, data
}

func (uc *invitationUsecase) GetByEmail(email string) *model.Invitation {
	result := <-uc.repo.FindByEmail(email)
	if result.Error != nil {
		debug.Println(result.Error)
		return new(model.Invitation)
	}
	data, _ := result.Data.(*model.Invitation)
	debug.PrintJSON(data)
	return data
}

func (uc *invitationUsecase) GetByName(name string) (int, []model.Invitation) {
	result := <-uc.repo.FindByName(name)
	if result.Error != nil {
		return 0, nil
	}
	data, _ := result.Data.([]model.Invitation)
	return result.Count, data
}

func (uc *invitationUsecase) Save(obj *model.Invitation) error {
	result := <-uc.repo.Save(obj)
	return result.Error
}
