package models

import (
	"gopkg.in/mgo.v2"
	"os"
)

type DB struct {
	*mgo.Database
}

type DBSession struct {
	*mgo.Session
}

func NewDB() (*DBSession, error) {
	db, err := mgo.Dial(os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}

	return &DBSession{db}, err
}

func (s *DBSession) EnsureIndexes() error {
	return nil
}
