package config

import (
	"fmt"
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

type databaseConn struct{}

// LoadDB mongo
func (d *databaseConn) LoadDB() *mgo.Database {
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	mongoHost := fmt.Sprintf("%s:%s", host, port)
	mongoSession, err := mgo.Dial(mongoHost)
	if err != nil {
		log.Fatal(err)
	}

	db := mongoSession.DB("wedding")

	// Init database collection, set unique index
	go func() {
		coll := db.C("invitations")
		index := mgo.Index{
			Key:    []string{"email"},
			Unique: true,
		}
		coll.EnsureIndex(index)
	}()

	return db
}
