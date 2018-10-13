package repository

import (
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
)

// Result repository model
type Result struct {
	Count int
	Data  interface{}
	Error error
}

// InvitationRepository abstract interface
type InvitationRepository interface {
	FindAll(offset, limit int) <-chan Result
	FindByWaNumber(waNumber string) <-chan Result
	FindByName(name string) <-chan Result
	CalculateCount(isAttend bool) <-chan Result
	Save(data *model.Invitation) <-chan Result
	RemoveByWaNumber(waNumber string) <-chan Result
}
