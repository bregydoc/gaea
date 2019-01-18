package gaea

import (
	"time"

	"github.com/oklog/ulid"
)

// AccountType is a kind of account
type AccountType string

// Anonymous is a Anonymous account kind
var Anonymous AccountType = "anonymous"

// Username is a Username account kind
var Username AccountType = "username"

// Phone is a Phone account kind
var Phone AccountType = "phone"

// Email is a Email account kind
var Email AccountType = "email"

// Google is a Google account kind
var Google AccountType = "google"

// Facebook is a Facebook account kind
var Facebook AccountType = "facebook"

// Twitter is a Twitter account kind
var Twitter AccountType = "twitter"

// Github is a Github account kind
var Github AccountType = "github"

// Account is a struct to represent the way how to person connect with our services
type Account struct {
	Person       ulid.ULID   `json:"person"`
	Type         AccountType `json:"type"`
	ID           string      `json:"id"`
	PasswordHash string      `json:"password"`
	SignUpAt     time.Time   `json:"sign_up_at"`
}
