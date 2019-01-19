package gaea

import (
	"time"

	"github.com/oklog/ulid"
)

// Token describe own token ticket
type Token struct {
	Owner       ulid.ULID       `json:"owner" bson:"owner"`
	Permissions map[string]byte `json:"permissions" bson:"permissions"`
	CreatedAt   time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" bson:"updated_at"`
	SignedBy    ulid.ULID       `json:"signed_by" bson:"signed_by"`
}
