package mongo

import (
	"context"
	"errors"

	"github.com/bregydoc/gaea"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/oklog/ulid"
)

// CreateAccount is for implement Gaea Repository interface
func (m *Repo) CreateAccount(c context.Context, account *gaea.Account) (*gaea.Account, error) {
	panic("unimplemented")
}

// GetAccountByMongoID returns an account find by mongo id ("_id")
func (m *Repo) GetAccountByMongoID(c context.Context, id string) (*gaea.Account, error) {
	a := new(gaea.Account)
	err := m.accountsCollection.FindOne(c, bson.M{"_id": id}).Decode(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// GetAccountByID is for implement Gaea Repository interface
func (m *Repo) GetAccountByID(c context.Context, id ulid.ULID) (*gaea.Account, error) {
	a := new(gaea.Account)
	err := m.accountsCollection.FindOne(c, bson.M{"id": id}).Decode(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// UpdateAccount is for implement Gaea Repository interface
func (m *Repo) UpdateAccount(c context.Context, account *gaea.Account) (*gaea.Account, error) {
	res, err := m.accountsCollection.UpdateOne(c, bson.M{"id": account.ID}, account)
	if err != nil {
		return nil, err
	}
	return m.GetAccountByMongoID(c, res.UpsertedID.(string))
}

// DeleteAccount is for implement Gaea Repository interface
func (m *Repo) DeleteAccount(c context.Context, account *gaea.Account) (*gaea.Account, error) {
	res, err := m.accountsCollection.DeleteOne(c, bson.M{"id": account.ID})
	if err != nil {
		return nil, err
	}

	if res.DeletedCount == 0 {
		// TODO: Put this error into global detection
		return nil, errors.New("delete failed")
	}

	return account, nil
}
