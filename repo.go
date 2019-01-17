package gaea

import (
	"context"
	"github.com/oklog/ulid"
)

type MinimalPersonInformation struct {
	Name     string   `json:"name"`
	Account  *Account `json:"account"`
	Password string   `json:"password"`
}

type Repository interface {
	Connect(ctx context.Context) error
	SignUp(ctx context.Context, info *MinimalPersonInformation) (*Person, error)                                      // Register
	SignIn(ctx context.Context, personID ulid.ULID, account *Account) (*Person, *Session, error)                      // Create session
	SignOut(ctx context.Context, personID ulid.ULID, sessionID ulid.ULID) (*Person, error)                            // Logout Session
	UpdateAccount(ctx context.Context, personID ulid.ULID, oldAccount *Account, newAccount *Account) (*Person, error) // e.g. reset password
	UpdatePersonInformation(ctx context.Context, personID ulid.ULID, newInfo *Person) (*Person, error)                // Update profile
}
