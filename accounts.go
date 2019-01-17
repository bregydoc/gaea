package gaea

import (
	"github.com/oklog/ulid"
	"time"
)

type AccountType string

var Anonymous AccountType = "anonymous"
var Username AccountType = "username"
var Phone AccountType = "phone"
var Email AccountType = "email"
var Google AccountType = "google"
var Facebook AccountType = "facebook"
var Twitter AccountType = "twitter"
var Github AccountType = "github"

type Account struct {
	Person       ulid.ULID   `json:"person"`
	Type         AccountType `json:"type"`
	ID           string      `json:"id"`
	PasswordHash string      `json:"password"`
	SignUpAt     time.Time   `json:"sign_up_at"`
}
