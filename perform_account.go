package gaea

import (
	"net/mail"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var phoneRegex = regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)

// GenerateAndEncodeAccount generate an account based on this basic information
func GenerateAndEncodeAccount(account *Account, pass string) (*Account, error) {
	newAccount := new(Account)
	switch account.Type {
	case Email:
		a, err := mail.ParseAddress(account.ID)
		if err != nil {
			return nil, err
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		newAccount.ID = a.Address
		newAccount.PasswordHash = string(hash)
		newAccount.SignUpAt = time.Now()
		newAccount.Type = Email

		return newAccount, nil

	case Phone:
		if !phoneRegex.Match([]byte(account.ID)) {
			return nil, ErrInvalidPhoneNumber
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		newAccount.ID = account.ID
		newAccount.PasswordHash = string(hash)
		newAccount.SignUpAt = time.Now()
		newAccount.Type = Phone

		return newAccount, nil

	case Username:
		hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		newAccount.ID = account.ID
		newAccount.PasswordHash = string(hash)
		newAccount.SignUpAt = time.Now()
		newAccount.Type = Username
		return newAccount, nil
	case Google:
		return nil, ErrUnimplementedError
	default:
		return nil, ErrAccountTypeNotExist
	}
}

// CheckIfPasswordIsCorrect only compare a original password with hashed password using bcrypt
func CheckIfPasswordIsCorrect(account *Account, password string) (*Account, error) {
	hash := []byte(account.PasswordHash)
	pass := []byte(password)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return account, err
}
