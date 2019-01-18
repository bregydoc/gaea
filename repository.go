package gaea

import (
	"context"

	"github.com/oklog/ulid"
)

// Repository is an interface to determinated repository
type Repository interface {
	CreatePerson(c context.Context, person *Person) (*Person, error)
	GetPersonByID(c context.Context, id ulid.ULID) (*Person, error)
	GetPersonByAccount(c context.Context, account *Account) (*Person, error)
	GetPersonBySession(c context.Context, account *Session) (*Person, error)
	UpdatePerson(c context.Context, person *Person) (*Person, error)
	DeletePerson(c context.Context, person *Person) (*Person, error)

	CreateAccount(c context.Context, account *Account) (*Account, error)
	GetAccountByID(c context.Context, id ulid.ULID) (*Account, error)
	UpdateAccount(c context.Context, account *Account) (*Account, error)
	DeleteAccount(c context.Context, account *Account) (*Account, error)

	CreateSession(c context.Context, session *Session) (*Session, error)
	GetSessionByID(c context.Context, id ulid.ULID) (*Session, error)
	UpdateSession(c context.Context, session *Session) (*Session, error)
	DeleteSession(c context.Context, session *Session) (*Session, error)
}
