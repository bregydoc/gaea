package gaea

import (
	"time"

	"github.com/oklog/ulid"
)

// AccountType is a kind of account
type AccountType string

// Anonymous is a kind of account
var Anonymous AccountType = "anonymous"

// Username is a kind of account
var Username AccountType = "username"

// Phone is a kind of account
var Phone AccountType = "phone"

// Email is a kind of account
var Email AccountType = "email"

// Google is a kind of account
var Google AccountType = "google"

// Facebook is a kind of account
var Facebook AccountType = "facebook"

// Twitter is a kind of account
var Twitter AccountType = "twitter"

// Github is a kind of account
var Github AccountType = "github"

// Account is a struct to represent the way how to person connect with our services
type Account struct {
	Person       ulid.ULID   `json:"person"`
	Type         AccountType `json:"type"`
	ID           string      `json:"id"`
	PasswordHash string      `json:"password"`
	SignUpAt     time.Time   `json:"sign_up_at"`
}
