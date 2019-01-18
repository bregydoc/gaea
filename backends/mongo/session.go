package mongo

import (
	"context"

	"github.com/bregydoc/gaea"
	"github.com/oklog/ulid"
)

// CreateSession implemented the Gaea Repository interface
func (m *Repo) CreateSession(c context.Context, session *gaea.Session) (*gaea.Session, error) {
	panic("unimplemented")
}

// GetSessionByID implemented the Gaea Repository interface
func (m *Repo) GetSessionByID(c context.Context, id ulid.ULID) (*gaea.Session, error) {
	panic("unimplemented")
}

// UpdateSession implemented the Gaea Repository interface
func (m *Repo) UpdateSession(c context.Context, session *gaea.Session) (*gaea.Session, error) {
	panic("unimplemented")
}

// DeleteSession implemented the Gaea Repository interface
func (m *Repo) DeleteSession(c context.Context, session *gaea.Session) (*gaea.Session, error) {
	panic("unimplemented")
}
