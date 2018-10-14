package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID       bson.ObjectId `bson:"_id" db:"id" json:"id,omitempty"`
	Name     string        `bson:"name,omitempty" db:"name" json:"name,omitempty"`
	Username string        `bson:"username,omitempty" db:"username" json:"username,omitempty"`
	Password string        `bson:"password,omitempty" db:"password" json:"password,omitempty"`
	Role     *Role         `bson:"role" json:"role,omitempty"`
}

type Role struct {
	ID       bson.ObjectId `bson:"_id" db:"id" json:"id,omitempty"`
	StringID string        `bson:"string_id" db:"string_id" json:"string_id,omitempty"`
	Name     string        `bson:"name" db:"name" json:"name,omitempty"`
}
