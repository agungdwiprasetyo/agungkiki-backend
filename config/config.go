package config

import (
	"crypto/rsa"

	"gopkg.in/mgo.v2"
)

// Config abstraction
type Config interface {
	LoadDB() *mgo.Database

	LoadPublicKey() *rsa.PublicKey
	LoadPrivateKey() *rsa.PrivateKey
}

// New init config
func New() Config {
	var conf struct {
		databaseConn
		rsaKey
	}
	return &conf
}
