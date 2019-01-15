package gaea

import (
	"github.com/oklog/ulid"
	"time"
)

type MID string

type Date struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

type EmailAddress struct {
	Address     string   `json:"address"`
	DisplayName string   `json:"display_name"`
	Type        string   `json:"type"`
	Extras      []string `json:"extras"`
	Valid       bool     `json:"valid"`
}

type PhoneNumber struct {
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
	Valid       bool   `json:"valid"`
}

type IdentityDocument struct {
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
	Document    string `json:"document"`
	Valid       bool   `json:"valid"`
}

type Person struct {
	ID       ulid.ULID `json:"id"`
	MID      MID       `json:"mid"`
	Name     string    `json:"name"`
	Photo    string    `json:"photo"`
	Age      int32     `json:"age"`
	Birthday *Date     `json:"birthday"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	IdentityDocument      *IdentityDocument   `json:"identity_document"`
	ExtraIdentityDocument []*IdentityDocument `json:"extra_identity_document"`

	EmailAccount      *EmailAddress   `json:"main_email"`
	ExtraEmailAccount []*EmailAddress `json:"extra_email"`

	PhoneNumber      *PhoneNumber `json:"phone_number"`
	ExtraPhoneNumber *PhoneNumber `json:"extra_phone_number"`

	Groups   []Group    `json:"group"`
	Accounts []*Account `json:"accounts"`

	ActiveSessions []*Session `json:"active_sessions"`

	Logs []*Log `json:"logs"`
}
