package repository

import (
	"time"

	"github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type invitationRepo struct {
	db *mgo.Database
}

// NewInvitationRepository create new repository
func NewInvitationRepository(repo *Repository) InvitationRepository {
	ir := new(invitationRepo)
	ir.db = repo.db
	return ir
}

func (r *invitationRepo) FindAll(offset, limit int, isAttend *bool) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		var invitations []model.Invitation
		q := bson.M{}
		if isAttend != nil {
			q = bson.M{"is_attend": *isAttend}
		}
		query := r.db.C("invitations").Find(q).Sort("-created")
		count, _ := query.Count()
		if err := query.Skip(offset).Limit(limit).All(&invitations); err != nil {
			output <- Result{Error: err}
			return
		}

		output <- Result{Count: count, Data: invitations}
	}()

	return output
}

func (r *invitationRepo) FindByWaNumber(waNumber string) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		var invitation model.Invitation
		query := bson.M{"wa_number": waNumber}
		if err := r.db.C("invitations").Find(query).One(&invitation); err != nil {
			output <- Result{Error: err}
			return
		}

		output <- Result{Data: &invitation}
	}()

	return output
}

func (r *invitationRepo) FindByName(name string) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		var invitations []model.Invitation
		query := r.db.C("invitations").Find(bson.M{"name": bson.M{"$regex": name}})
		count, _ := query.Count()
		if err := query.All(&invitations); err != nil {
			output <- Result{Error: err}
			return
		}

		output <- Result{Count: count, Data: invitations}
	}()

	return output
}

func (r *invitationRepo) CalculateCount(isAttend bool) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		query := r.db.C("invitations").Find(bson.M{"is_attend": isAttend})
		count, err := query.Count()
		if err != nil {
			output <- Result{Error: err}
			return
		}

		output <- Result{Count: count}
	}()

	return output
}

func (r *invitationRepo) Save(obj *model.Invitation) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		loc, _ := time.LoadLocation("Asia/Jakarta")

		obj.ID = bson.NewObjectId()
		obj.CreatedAt = time.Now().In(loc)
		if err := r.db.C("invitations").Insert(obj); err != nil {
			output <- Result{Error: err}
			return
		}

		output <- Result{Data: obj}
	}()

	return output
}

func (r *invitationRepo) RemoveByWaNumber(number string) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)

		if err := r.db.C("invitations").Remove(bson.M{"wa_number": number}); err != nil {
			output <- Result{Error: err}
		}
	}()

	return output
}
