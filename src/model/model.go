package model

import (
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// Invitation model
type Invitation struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Message  string        `json:"message" bson:"message"`
	IsAttend bool          `json:"isAttend" bson:"is_attend"`
}

var InvitationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Invitation",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"isAttend": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
