package gaea

import (
	"net"
	"time"

	"github.com/oklog/ulid"
)

// Session define a session person (links person with device)
type Session struct {
	ID          ulid.ULID         `json:"id" bson:"id"`
	CreatedAt   time.Time         `json:"created_at" bson:"created_at"`
	Begin       time.Time         `json:"begin" bson:"begin"`
	Duration    time.Duration     `json:"duration" bson:"duration"`
	SessionType AccountType       `json:"session_type" bson:"session_type"`
	WithAccount *Account          `json:"with_account" bson:"with_account"`
	Token       string            `json:"token" bson:"token"`
	ClientIP    net.IP            `json:"client_ip" bson:"client_ip"`
	Metadata    map[string]string `json:"metadata" bson:"metadata"`
}

// NewSession returns a new session
func NewSession(account *Account, duration time.Duration, clientIP net.IP) (*Session, error) {
	s := &Session{
		CreatedAt:   time.Now(),
		ID:          newULID(),
		Duration:    duration,
		Begin:       time.Now(),
		ClientIP:    clientIP,
		Metadata:    map[string]string{},
		SessionType: account.Type,
		WithAccount: account,
		Token:       "", // TODO: Generate powerful tokens
	}
	// TODO: Improve this function
	return s, nil
}
