package usecase

import (
	"fmt"

	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"github.com/agungdwiprasetyo/agungkiki-backend/src/repository"
	tokenModule "github.com/agungdwiprasetyo/agungkiki-backend/token"
	"github.com/agungdwiprasetyo/go-utils"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

// UcResult usecase common result
type UcResult struct {
	Count int
	Data  interface{}
	Error error
}

type invitationUsecase struct {
	token          tokenModule.Token
	invitationRepo repository.InvitationRepository
	userRepo       repository.UserRepository
}

// NewInvitationUsecase create new usecase
func NewInvitationUsecase(token tokenModule.Token, repo *repository.Repository) InvitationUsecase {
	uc := new(invitationUsecase)
	uc.token = token
	uc.invitationRepo = repository.NewInvitationRepository(repo)
	uc.userRepo = repository.NewUserRepository(repo)
	return uc
}

func (uc *invitationUsecase) GetAll(params *model.AllInvitationParam) (result UcResult) {
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 10
	}
	params.Offset = (params.Page - 1) * params.Limit
	res := <-uc.invitationRepo.FindAll(params.Offset, params.Limit, params.IsAttend)
	if res.Error != nil {
		result.Error = res.Error
		return
	}
	result.Count = res.Count
	result.Data = res.Data
	return
}

func (uc *invitationUsecase) GetByWaNumber(waNumber string) *model.Invitation {
	result := <-uc.invitationRepo.FindByWaNumber(waNumber)
	if result.Error != nil {
		debug.Println(result.Error)
		return new(model.Invitation)
	}
	data, _ := result.Data.(*model.Invitation)
	debug.PrintJSON(data)
	return data
}

func (uc *invitationUsecase) GetByName(name string) (int, []model.Invitation) {
	result := <-uc.invitationRepo.FindByName(name)
	if result.Error != nil {
		return 0, nil
	}
	data, _ := result.Data.([]model.Invitation)
	return result.Count, data
}

func (uc *invitationUsecase) GetCount(isAttend bool) (int, error) {
	result := <-uc.invitationRepo.CalculateCount(isAttend)
	if result.Error != nil {
		return 0, nil
	}
	return result.Count, result.Error
}

func (uc *invitationUsecase) GetEvent() (result UcResult) {
	res := <-uc.invitationRepo.FindEvents()
	if res.Error != nil {
		result.Error = res.Error
		return
	}

	result.Data = res.Data
	return
}

func (uc *invitationUsecase) Save(obj *model.Invitation) error {
	result := <-uc.invitationRepo.Save(obj)
	return result.Error
}

func (uc *invitationUsecase) Remove(numbers []string) error {
	for _, number := range numbers {
		res := <-uc.invitationRepo.RemoveByWaNumber(number)
		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}

func (uc *invitationUsecase) UserLogin(username, password string) (result UcResult) {
	res := <-uc.userRepo.FindByUsername(username)
	if res.Error != nil {
		result.Error = fmt.Errorf("Username or Password is invalid")
		return
	}
	userData, _ := res.Data.(*model.User)

	hashedPassword := utils.ComputeHmac256(password, "mantul")
	if userData.Password != hashedPassword {
		result.Error = fmt.Errorf("Username or Password is invalid")
		return
	}

	tokenClaim := tokenModule.NewClaim(userData)
	tokenString, err := uc.token.Generate(tokenClaim)
	if err != nil {
		return UcResult{Error: err}
	}

	var data struct {
		Token string      `json:"token"`
		User  *model.User `json:"user"`
	}
	data.Token = tokenString
	data.User = userData

	return UcResult{Data: data}
}

func (uc *invitationUsecase) SaveUser(dataUser *model.User) (result UcResult) {
	dataUser.Password = utils.ComputeHmac256(dataUser.Password, "mantul")
	res := <-uc.userRepo.Insert(dataUser)
	if res.Error != nil {
		result.Error = res.Error
		return
	}
	result.Data = res.Data
	return
}
