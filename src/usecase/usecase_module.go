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

func (uc *invitationUsecase) GetByWaNumber(waNumber string) *model.Invitation {
	result := <-uc.repo.FindByWaNumber(waNumber)
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

func (uc *invitationUsecase) Remove(numbers []string) error {
	for _, number := range numbers {
		res := <-uc.repo.RemoveByWaNumber(number)
		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}
