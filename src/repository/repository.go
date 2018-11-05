package repository

import (
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"gopkg.in/mgo.v2"
)

// Result repository model
type Result struct {
	Count int
	Data  interface{}
	Error error
}

// InvitationRepository abstract interface
type InvitationRepository interface {
	FindAll(offset, limit int, isAttend *bool) <-chan Result
	FindByWaNumber(waNumber string) <-chan Result
	FindByName(name string) <-chan Result
	CalculateCount(isAttend bool) <-chan Result
	FindEvents() <-chan Result
	Save(data *model.Invitation) <-chan Result
	SaveEvent(data *model.Event) <-chan Result
	RemoveByWaNumber(waNumber string) <-chan Result
}

// UserRepository abstract interface
type UserRepository interface {
	FindByUsername(username string) <-chan Result
	Insert(dataUser *model.User) <-chan Result
}

type Repository struct {
	db *mgo.Database
}

func NewRepository(db *mgo.Database) *Repository {
	return &Repository{db}
}
