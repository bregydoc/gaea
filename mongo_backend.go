package gaea

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

type MongoRepo struct {
	client *mongo.Client
	config *Config

	personCollection   *mongo.Collection
	sessionsCollection *mongo.Collection
}

func NewMongoRepo(config *Config) (*MongoRepo, error) {
	client, err := mongo.NewClient(config.UriConnection)
	if err != nil {
		return nil, err
	}

	p := client.Database(config.UsersDBName).Collection(config.UsersDBName)
	s := client.Database(config.SessionsDBName).Collection(config.SessionsDBName)

	return &MongoRepo{
		config:             config,
		client:             client,
		personCollection:   p,
		sessionsCollection: s,
	}, nil
}

func (m *MongoRepo) Connect(ctx context.Context) error {
	return m.client.Connect(ctx)
}

func (m *MongoRepo) SignUp(ctx context.Context, info *MinimalPersonInformation) (*Person, error) {
	// Checkin Account Person
	person, err := ModelPersonWithMinimalInformation(info)
	if err != nil {
		return nil, err
	}
	person.UpdatedAt = time.Now()

	res, err := m.personCollection.InsertOne(ctx, person)
	if err != nil {
		return nil, err
	}

	savedPerson := new(Person)

	err = m.personCollection.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(savedPerson)

	if err != nil {
		return nil, err
	}

	return savedPerson, nil
}
