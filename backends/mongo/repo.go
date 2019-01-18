package mongo

import (
	"context"

	"github.com/bregydoc/gaea"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// Repo is a gaea repository
type Repo struct {
	client *mongo.Client
	config *gaea.Config

	peopleCollection   *mongo.Collection
	sessionsCollection *mongo.Collection
	accountsCollection *mongo.Collection
}

// NewRepo create new mongo repository
func NewRepo(config *gaea.Config) (*Repo, error) {
	client, err := mongo.NewClient(config.UriConnection)
	if err != nil {
		return nil, err
	}

	// TODO: Fix this, please
	p := client.Database(config.PeopleDBName).Collection(config.PeopleDBName)
	a := client.Database(config.PeopleDBName).Collection(config.AccountsDBName)
	s := client.Database(config.SessionsDBName).Collection(config.SessionsDBName)

	return &Repo{
		config:             config,
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
