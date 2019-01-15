package gaea

import (
	"github.com/oklog/ulid"
	"net"
	"time"
)

type Session struct {
	CreatedAt   *time.Time        `json:"created_at"`
	Begin       *time.Time        `json:"begin"`
	Duration    *time.Duration    `json:"duration"`
	SessionType *AccountType      `json:"session_type"`
	ID          ulid.ULID         `json:"code"`
	Token       string            `json:"token"`
	ClientIP    net.IP            `json:"client_ip"`
	Metadata    map[string]string `json:"metadata"`
}
