package repository

import (
	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userRepo struct {
	db *mgo.Database
}

// NewUserRepository create new repository
func NewUserRepository(repo *Repository) UserRepository {
	ur := new(userRepo)
	ur.db = repo.db
	return ur
}

func (r *userRepo) FindByUsername(username string) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		var userData model.User
		query := bson.M{"username": username}
		if err := r.db.C("users").Find(query).One(&userData); err != nil {
			output <- Result{Error: err}
			return
		}

		output <- Result{Data: &userData}
	}()

	return output
}

func (r *userRepo) Insert(userData *model.User) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		userData.ID = bson.NewObjectId()
		if userData.Role != nil {
			userData.Role.ID = bson.NewObjectId()
		}
		if err := r.db.C("users").Insert(userData); err != nil {
			output <- Result{Error: err}
			return
		}

		output <- Result{Data: userData.ID}
	}()

	return output
}
