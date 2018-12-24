package model

import (
	"time"

	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// Visitor model
type Visitor struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	IPAddress string        `json:"ip_address,omitempty" bson:"ip_address"`
	Query     string        `json:"query,omitempty" bson:"query"`
	Datetime  time.Time     `json:"datetime,omitempty" bson:"datetime"`
}

func (v *Visitor) MakeObject() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Invitation",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"ip_address": &graphql.Field{
				Type: graphql.String,
			},
			"query": &graphql.Field{
				Type: graphql.String,
			},
			"datetime": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}
