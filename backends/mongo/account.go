package mongo

import (
	"context"

	"github.com/bregydoc/gaea"
	"github.com/oklog/ulid"
)

// CreateAccount is for implement Gaea Repository interface
func (m *Repo) CreateAccount(c context.Context, account *gaea.Account) (*gaea.Account, error) {
	panic("unimplemented")
}

// GetAccountByID is for implement Gaea Repository interface
func (m *Repo) GetAccountByID(c context.Context, id ulid.ULID) (*gaea.Account, error) {
	panic("unimplemented")
}

// UpdateAccount is for implement Gaea Repository interface
func (m *Repo) UpdateAccount(c context.Context, account *gaea.Account) (*gaea.Account, error) {
	panic("unimplemented")
}

// DeleteAccount is for implement Gaea Repository interface
func (m *Repo) DeleteAccount(c context.Context, account *gaea.Account) (*gaea.Account, error) {
	panic("unimplemented")
}
