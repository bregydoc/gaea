package gaea

import (
	"context"

	"github.com/oklog/ulid"
)

// MinimalPersonInformation is the minimal information to indentify a person
type MinimalPersonInformation struct {
	Name     string   `json:"name"`
	Account  *Account `json:"account"`
	Password string   `json:"password"`
}

// Gaea is the interface to mock how to works Gaea
type Gaea interface {
	SignUp(c context.Context, info *MinimalPersonInformation) (*Person, error)                                      // Register
	SignIn(c context.Context, personID ulid.ULID, account *Account, password string) (*Person, *Session, error)     // Create session
	SignOut(c context.Context, personID ulid.ULID, sessionID ulid.ULID) (*Person, error)                            // Logout Session
	UpdateAccount(c context.Context, personID ulid.ULID, oldAccount *Account, newAccount *Account) (*Person, error) // e.g. reset password
	UpdatePersonInformation(c context.Context, personID ulid.ULID, newInfo *Person) (*Person, error)                // Update profile
}
