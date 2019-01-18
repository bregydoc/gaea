package mongo

import (
	"context"

	"github.com/bregydoc/gaea"
	"github.com/oklog/ulid"
)

// CreatePerson is for implemet Gaea Repository interface
func (m *Repo) CreatePerson(c context.Context, person *gaea.Person) (*gaea.Person, error) {
	panic("unimplemented")
}

// GetPersonByID is for implemet Gaea Repository interface
func (m *Repo) GetPersonByID(c context.Context, id ulid.ULID) (*gaea.Person, error) {
	panic("unimplemented")
}

// GetPersonByAccount is for implemet Gaea Repository interface
func (m *Repo) GetPersonByAccount(c context.Context, account *gaea.Account) (*gaea.Person, error) {
	panic("unimplemented")
}

// GetPersonBySession is for implemet Gaea Repository interface
func (m *Repo) GetPersonBySession(c context.Context, account *gaea.Session) (*gaea.Person, error) {
	panic("unimplemented")
}

// UpdatePerson is for implemet Gaea Repository interface
func (m *Repo) UpdatePerson(c context.Context, person *gaea.Person) (*gaea.Person, error) {
	panic("unimplemented")
}

// DeletePerson is for implemet Gaea Repository interface
func (m *Repo) DeletePerson(c context.Context, person *gaea.Person) (*gaea.Person, error) {
	panic("unimplemented")
}
