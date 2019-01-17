package gaea

import (
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"regexp"
	"time"
)

var phoneRegex = regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)

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
			return nil, invalidPhoneNumber
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
		return nil, unimplementedError
	default:
		return nil, accountTypeNotExist
	}
}

func CheckIfPasswordIsCorrect(account *Account, password string) (*Account, error) {
	hash := []byte(account.PasswordHash)
	pass := []byte(password)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return account, err
}
