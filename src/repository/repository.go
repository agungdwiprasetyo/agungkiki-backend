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
	FindAll() <-chan Result
	FindByEmail(email string) <-chan Result
	FindByName(name string) <-chan Result
	Save(data *model.Invitation) <-chan Result
}
