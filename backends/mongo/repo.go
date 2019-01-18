package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// Repo is a gaea repository
type Repo struct {
	client *mongo.Client

	peopleCollection   *mongo.Collection
	sessionsCollection *mongo.Collection
	accountsCollection *mongo.Collection
}

// NewRepo create new mongo repository
func NewRepo(uri, db, peopleCollection, accountsCollection, sessionsCollection string) (*Repo, error) {
	client, err := mongo.NewClient(uri)
	if err != nil {
		return nil, err
	}

	p := client.Database(db).Collection(peopleCollection)
	a := client.Database(db).Collection(accountsCollection)
	s := client.Database(db).Collection(sessionsCollection)

	return &Repo{
		client:             client,
		peopleCollection:   p,
		sessionsCollection: s,
		accountsCollection: a,
	}, nil

}

// Connect implement and lauch repository
func (m *Repo) Connect(ctx context.Context) error {
	err := m.client.Connect(ctx)
	if err != nil {
		return err
	}

	unique := true

	_, err = m.accountsCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"id": 1},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})

	_, err = m.peopleCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"id": 1},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})

	if err != nil {
		return err
	}

	return err
}
