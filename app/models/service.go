package models

import (
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	ID   string        `bson:"id" json:"id"`
	mID  bson.ObjectId `bson:"_id" json:"_id"`
	Name string        `bson:"name" json:"name"`
}

func (m *DB) NewService(name string) (*Service, error) {
	s := Service{
		ID:   uuid.NewV4().String(),
		mID:  bson.NewObjectId(),
		Name: name,
	}

	err := m.C("service").Insert(&s)
	return &s, err
}

func (m *DB) GetServiceByName(name string) (*Service, error) {
	var s Service
	err := m.C("service").Find(bson.M{"name": name}).One(&s)
	return &s, err
}

func (m *DB) GetServicesByPage(page, pageSize int) ([]*Service, error) {
	var s []*Service
	err := m.C("service").Find(nil).Sort("-_id").Skip(pageSize * (page - 1)).Limit(pageSize).All(&s)
	return s, err
}

func (m *DB) DeleteServiceByName(name string) error {
	return m.C("service").Remove(bson.M{"name": name})
}
