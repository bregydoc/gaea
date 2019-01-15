package gaea

import (
	"time"
)

func ModelPersonWithMinimalInformation(info *MinimalPersonInformation) (*Person, error) {
	// Creating person
	account, err := GenerateAndEncodeAccount(info.Account, info.Password)
	if err != nil {
		return nil, err
	}

	person := &Person{
		ID:        newULID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      info.Name,
		Logs: []*Log{{
			At:      time.Now(),
			Type:    Mutation,
			Content: "person profile created",
		},
		},
		Groups: []Group{Customer},
	}

	switch account.Type {
	case Email:
		person.EmailAccount = &EmailAddress{
			Type:        "main",
			Address:     account.ID,
			DisplayName: "Principal",
			Valid:       false,
		}
		person.Accounts = []*Account{
			account,
		}
		break
	case Phone:
		phone, err := GetPhoneNumberFromString(account.ID)
		if err != nil {
			return nil, err
		}

		person.PhoneNumber = &PhoneNumber{
			Type:        "main",
			CountryCode: phone.CountryCode,
			DisplayName: "Personal",
			Number:      phone.Number,
			Valid:       false,
		}
		person.Accounts = []*Account{
			account,
		}
		break
	case Username:

		person.Accounts = []*Account{
			account,
		}
		break
	case Google:
		person.Accounts = []*Account{
			account,
		}
		break
	default:
		return nil, accountTypeNotExist
	}

	return person, nil
}
