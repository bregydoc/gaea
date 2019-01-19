package mongo

import (
	"context"
	"errors"

	"github.com/bregydoc/gaea"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/oklog/ulid"
)

// CreatePerson is for implemet Gaea Repository interface
func (m *Repo) CreatePerson(c context.Context, person *gaea.Person) (*gaea.Person, error) {
	r, err := m.peopleCollection.InsertOne(c, person)
	if err != nil {
		return nil, err
	}

	return m.GetPersonByMongoID(c, r.InsertedID)
}

// GetPersonByMongoID returns a person by mongo ID ("_id")
func (m *Repo) GetPersonByMongoID(c context.Context, id interface{}) (*gaea.Person, error) {
	p := new(gaea.Person)
	err := m.peopleCollection.FindOne(c, bson.M{"_id": id}).Decode(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// GetPersonByID is for implemet Gaea Repository interface
func (m *Repo) GetPersonByID(c context.Context, id ulid.ULID) (*gaea.Person, error) {
	p := new(gaea.Person)
	err := m.peopleCollection.FindOne(c, bson.M{"id": id}).Decode(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// GetPersonByAccount is for implemet Gaea Repository interface
func (m *Repo) GetPersonByAccount(c context.Context, account *gaea.Account) (*gaea.Person, error) {
	return m.GetPersonByID(c, account.Person)
}

// GetPersonBySession is for implemet Gaea Repository interface
func (m *Repo) GetPersonBySession(c context.Context, session *gaea.Session) (*gaea.Person, error) {
	return m.GetPersonByID(c, session.WithAccount.Person)
}

// UpdatePerson is for implemet Gaea Repository interface
func (m *Repo) UpdatePerson(c context.Context, person *gaea.Person) (*gaea.Person, error) {
	res, err := m.peopleCollection.UpdateOne(c, bson.M{"id": person.ID}, person)
	if err != nil {
		return nil, err
	}

	return m.GetPersonByMongoID(c, res.UpsertedID.(string))
}

// DeletePerson is for implemet Gaea Repository interface
func (m *Repo) DeletePerson(c context.Context, person *gaea.Person) (*gaea.Person, error) {
	res, err := m.peopleCollection.DeleteOne(c, bson.M{"id": person.ID})
	if err != nil {
		return nil, err
	}

	if res.DeletedCount == 0 {
		// TODO: Put this error into global detection
		return nil, errors.New("delete failed")
	}

	return person, nil
}
