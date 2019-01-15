package gaea

import (
	"github.com/oklog/ulid"
	"time"
)

type MID string

type Date struct {
	Day   int64 `json:"day" bson:"day"`
	Month int64 `json:"month" bson:"month"`
	Year  int64 `json:"year" bson:"year"`
}

type EmailAddress struct {
	Address     string   `json:"address" bson:"address"`
	DisplayName string   `json:"display_name" bson:"display_name"`
	Type        string   `json:"type" bson:"type"`
	Extras      []string `json:"extras" bson:"extras"`
	Valid       bool     `json:"valid" bson:"valid"`
}

type PhoneNumber struct {
	DisplayName string `json:"display_name" bson:"display_name"`
	Type        string `json:"type" bson:"type"`
	CountryCode string `json:"country_code" bson:"country_code"`
	Number      string `json:"number" bson:"number"`
	Valid       bool   `json:"valid" bson:"valid"`
}

type IdentityDocument struct {
	DisplayName string `json:"display_name" bson:"display_name"`
	Type        string `json:"type" bson:"type"`
	Document    string `json:"document" bson:"document"`
	Valid       bool   `json:"valid" bson:"valid"`
}

type Person struct {
	ID       ulid.ULID `json:"id" bson:"id"`
	MID      MID       `json:"mid" bson:"mid"`
	Name     string    `json:"name" bson:"name"`
	Photo    string    `json:"photo" bson:"photo"`
	Age      int32     `json:"age" bson:"age"`
	Birthday *Date     `json:"birthday" bson:"birthday"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

	IdentityDocument      *IdentityDocument   `json:"identity_document" bson:"identity_document"`
	ExtraIdentityDocument []*IdentityDocument `json:"extra_identity_document" bson:"extra_identity_document"`

	EmailAccount      *EmailAddress   `json:"main_email" bson:"main_email"`
	ExtraEmailAccount []*EmailAddress `json:"extra_email" bson:"extra_email"`

	PhoneNumber      *PhoneNumber `json:"phone_number" bson:"phone_number"`
	ExtraPhoneNumber *PhoneNumber `json:"extra_phone_number" bson:"extra_phone_number"`

	Groups   []Group    `json:"group" bson:"group"`
	Accounts []*Account `json:"accounts" bson:"accounts"`

	ActiveSessions []*Session `json:"active_sessions" bson:"active_sessions"`

	Logs []*Log `json:"logs" bson:"logs"`
}
