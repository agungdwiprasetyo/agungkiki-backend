package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Visitor model
type Visitor struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	IPAddress string        `json:"ip_address,omitempty" bson:"ip_address"`
	Query     string        `json:"query,omitempty" bson:"query"`
	Datetime  time.Time     `json:"datetime,omitempty" bson:"datetime"`
}
