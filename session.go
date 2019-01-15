package gaea

import (
	"github.com/oklog/ulid"
	"net"
	"time"
)

type Session struct {
	CreatedAt   *time.Time        `json:"created_at" bson:"created_at"`
	Begin       *time.Time        `json:"begin" bson:"begin"`
	Duration    *time.Duration    `json:"duration" bson:"duration"`
	SessionType *AccountType      `json:"session_type" bson:"session_type"`
	ID          ulid.ULID         `json:"code" bson:"code"`
	Token       string            `json:"token" bson:"token"`
	ClientIP    net.IP            `json:"client_ip" bson:"client_ip"`
	Metadata    map[string]string `json:"metadata" bson:"metadata"`
}
