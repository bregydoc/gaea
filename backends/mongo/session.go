package mongo

import (
	"context"
	"errors"

	"github.com/bregydoc/gaea"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/oklog/ulid"
)

// CreateSession implemented the Gaea Repository interface
func (m *Repo) CreateSession(c context.Context, session *gaea.Session) (*gaea.Session, error) {
	res, err := m.sessionsCollection.InsertOne(c, session)
	if err != nil {
		return nil, err
	}
	return m.GetSessionByMongoID(c, res.InsertedID)
}

// GetSessionByMongoID returns a session by Mongo id ("_id")
func (m *Repo) GetSessionByMongoID(c context.Context, id interface{}) (*gaea.Session, error) {
	s := new(gaea.Session)
	err := m.sessionsCollection.FindOne(c, bson.M{"_id": id}).Decode(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetSessionByID implemented the Gaea Repository interface
func (m *Repo) GetSessionByID(c context.Context, id ulid.ULID) (*gaea.Session, error) {
	s := new(gaea.Session)
	err := m.sessionsCollection.FindOne(c, bson.M{"id": id}).Decode(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// UpdateSession implemented the Gaea Repository interface
func (m *Repo) UpdateSession(c context.Context, session *gaea.Session) (*gaea.Session, error) {
	res, err := m.sessionsCollection.UpdateOne(c, bson.M{"id": session.ID}, session)
	if err != nil {
		return nil, err
	}
	return m.GetSessionByMongoID(c, res.UpsertedID.(string))
}

// DeleteSession implemented the Gaea Repository interface
func (m *Repo) DeleteSession(c context.Context, session *gaea.Session) (*gaea.Session, error) {
	res, err := m.sessionsCollection.DeleteOne(c, bson.M{"id": session.ID})
	if err != nil {
		return nil, err
	}

	if res.DeletedCount == 0 {
		// TODO: Put this error into global detection
		return nil, errors.New("delete failed")
	}

	return session, nil
}
