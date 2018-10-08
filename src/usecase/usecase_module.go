package usecase

import (
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/repository"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

type invitationUsecase struct {
	repo repository.InvitationRepository
}

func (uc *invitationUsecase) GetAll(offset, limit int) (int, []model.Invitation) {
	offset = (offset - 1) * limit
	result := <-uc.repo.FindAll(offset, limit)
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

func (uc *invitationUsecase) GetCount(isAttend bool) (int, error) {
	result := <-uc.repo.CalculateCount(isAttend)
	if result.Error != nil {
		return 0, nil
	}
	return result.Count, result.Error
}

func (uc *invitationUsecase) Save(obj *model.Invitation) error {
	result := <-uc.repo.Save(obj)
	return result.Error
}

func (uc *invitationUsecase) Remove(emails []string) error {
	for _, email := range emails {
		res := <-uc.repo.RemoveByEmail(email)
		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}
