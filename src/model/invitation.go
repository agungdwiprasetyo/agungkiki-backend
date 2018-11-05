package model

import (
	"time"

	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// Invitation model
type Invitation struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	WaNumber  string        `json:"waNumber" bson:"wa_number"`
	Message   string        `json:"message" bson:"message"`
	IsAttend  bool          `json:"isAttend" bson:"is_attend"`
	CreatedAt time.Time     `json:"created" bson:"created"`
}

// Event model
type Event struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Code      string        `json:"code" bson:"code"`
	Date      string        `bson:"date" json:"date"`
	Ceremony  string        `bson:"ceremony" json:"ceremony"`
	Reception string        `bson:"reception" json:"reception"`
	Address   string        `bson:"address" json:"address"`
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
		"waNumber": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"isAttend": &graphql.Field{
			Type: graphql.Boolean,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
