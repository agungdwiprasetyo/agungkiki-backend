package config

import (
	"gopkg.in/mgo.v2"
)

// Config abstraction
type Config interface {
	LoadDB() *mgo.Database
}

// New init config
func New() Config {
	var conf struct {
		databaseConn
	}
	return &conf
}
