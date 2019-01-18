package mongo

type MongoRepo struct {
	// client *mongo.Client
	// config *Config

	// peopleCollection   *mongo.Collection
	// sessionsCollection *mongo.Collection
	// accountsCollection *mongo.Collection
}

// func NewMongoRepo(config *Config) (*MongoRepo, error) {
// 	client, err := mongo.NewClient(config.UriConnection)
// 	if err != nil {
// 		return nil, err
// 	}

// 	p := client.Database(config.PeopleDBName).Collection(config.PeopleDBName)
// 	a := client.Database(config.PeopleDBName).Collection(config.AccountsDBName)

// 	s := client.Database(config.SessionsDBName).Collection(config.SessionsDBName)

// 	return &MongoRepo{
// 		config:             config,
// 		client:             client,
// 		peopleCollection:   p,
// 		sessionsCollection: s,
// 		accountsCollection: a,
// 	}, nil
// }

// func (m *MongoRepo) Connect(ctx context.Context) error {
// 	err := m.client.Connect(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	unique := true

// 	_, err = m.accountsCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
// 		Keys: bson.M{"id": 1},
// 		Options: &options.IndexOptions{
// 			Unique: &unique,
// 		},
// 	})

// 	_, err = m.peopleCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
// 		Keys: bson.M{"id": 1},
// 		Options: &options.IndexOptions{
// 			Unique: &unique,
// 		},
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	return err
// }

// func (m *MongoRepo) SignUp(ctx context.Context, info *MinimalPersonInformation) (*Person, error) {
// 	// Checkin Account Person
// 	person, err := ModelPersonWithMinimalInformation(info)
// 	if err != nil {
// 		return nil, err
// 	}

// 	person.UpdatedAt = time.Now()

// 	if len(person.Accounts) != 1 {
// 		return nil, invalidSignUpPerson
// 	}

// 	// Saving account
// 	_, err = m.accountsCollection.InsertOne(ctx, person.Accounts[0])
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Save person process
// 	res, err := m.peopleCollection.InsertOne(ctx, person)
// 	if err != nil {
// 		return nil, err
// 	}

// 	savedPerson := new(Person)
// 	err = m.peopleCollection.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(savedPerson)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return savedPerson, nil
// }

// func (m *MongoRepo) SignIn(ctx context.Context, personID ulid.ULID, account *Account, password string) (*Person, *Session, error) {
// 	findAccount := new(Account)
// 	err := m.accountsCollection.FindOne(ctx, bson.M{"id": account.ID}).Decode(findAccount)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	person := new(Person)
// 	err = m.peopleCollection.FindOne(ctx, bson.M{"id": personID}).Decode(person)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	if person.ID != findAccount.Person {
// 		// TODO: I don't know if we need check this...
// 		// TODO: Fix if it's not necessary
// 		// TODO: Currently that returns a invalid account
// 		return nil, nil, accountPersonNoMatch
// 	}

// 	findAccount, err = CheckIfPasswordIsCorrect(findAccount, password)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	s, err := NewSession(findAccount, defaultSessionDuration, net.IPv4zero) // TODO: Change IP
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	if person.ActiveSessions == nil {
// 		person.ActiveSessions = []*Session{}
// 	}
// 	person.ActiveSessions = append(person.ActiveSessions, s)

// 	person.Logs = append(person.Logs, &Log{
// 		Type:    Info,
// 		At:      time.Now(),
// 		Content: "added new session to person",
// 		Fields: map[string]string{
// 			"session_id": s.ID.String(),
// 		},
// 	})

// 	up, err := m.peopleCollection.UpdateOne(ctx, bson.M{"id": person.ID}, person)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	// TODO: Check below
// 	err = m.peopleCollection.FindOne(ctx, bson.M{"_id": up.UpsertedID}).Decode(person)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	in, err := m.sessionsCollection.InsertOne(ctx, s)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	err = m.sessionsCollection.FindOne(ctx, bson.M{"_id": in.InsertedID}).Decode(s)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return person, s, nil
// }

// func (m *MongoRepo) SignOut(ctx context.Context, personID ulid.ULID, sessionID ulid.ULID) (*Person, error) {
// 	return &Person{}, nil
// }
