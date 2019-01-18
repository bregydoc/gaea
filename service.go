package gaea

import (
	"context"
	"net"
	"time"

	"github.com/oklog/ulid"
)

// Service ...
type Service struct {
	Repo Repository
}

// SignUp implemets the Gaea service
func (s *Service) SignUp(c context.Context, info *MinimalPersonInformation) (*Person, error) {
	// Checkin Account Person
	person, err := ModelPersonWithMinimalInformation(info)
	if err != nil {
		return nil, err
	}

	person.UpdatedAt = time.Now()

	if len(person.Accounts) != 1 {
		return nil, errInvalidSignUpPerson
	}

	// Saving account
	_, err = s.Repo.CreateAccount(c, person.Accounts[0])
	if err != nil {
		return nil, err
	}

	// Save person process
	newPerson, err := s.Repo.CreatePerson(c, person)
	if err != nil {
		return nil, err
	}

	return newPerson, nil
}

// SignIn implemets the Gaea service
func (s *Service) SignIn(c context.Context, personID ulid.ULID, account *Account, password string) (*Person, *Session, error) {

	findAccount, err := s.Repo.GetAccountByID(c, account.ID)
	if err != nil {
		return nil, nil, err
	}

	person, err := s.Repo.GetPersonByID(c, personID)
	if err != nil {
		return nil, nil, err
	}

	if person.ID != findAccount.Person {
		// TODO: I don't know if we need check this...
		// TODO: Fix if it's not necessary
		// TODO: Currently that returns a invalid account
		return nil, nil, errAccountPersonNoMatch
	}

	findAccount, err = CheckIfPasswordIsCorrect(findAccount, password)
	if err != nil {
		return nil, nil, err
	}

	session, err := NewSession(findAccount, defaultSessionDuration, net.IPv4zero) // TODO: Change IP
	if err != nil {
		return nil, nil, err
	}

	if person.ActiveSessions == nil {
		person.ActiveSessions = []*Session{}
	}
	person.ActiveSessions = append(person.ActiveSessions, session)

	person.Logs = append(person.Logs, &Log{
		Type:    Info,
		At:      time.Now(),
		Content: "added new session to person",
		Fields: map[string]string{
			"session_id": session.ID.String(),
		},
	})

	person, err = s.Repo.UpdatePerson(c, person)
	if err != nil {
		return nil, nil, err
	}

	session, err = s.Repo.CreateSession(c, session)
	if err != nil {
		return nil, nil, err
	}

	return person, session, nil
}

// SignOut implemets the Gaea service
func (s *Service) SignOut(c context.Context, personID ulid.ULID, sessionID ulid.ULID) (*Person, error) {
	return &Person{}, nil
}

// UpdateAccount implemets the Gaea service
func (s *Service) UpdateAccount(c context.Context, personID ulid.ULID, oldAccount *Account, newAccount *Account) (*Person, error) {
	return &Person{}, nil
}

// UpdatePersonInformation implemets the Gaea service
func (s *Service) UpdatePersonInformation(c context.Context, personID ulid.ULID, newInfo *Person) (*Person, error) {
	return &Person{}, nil
}
