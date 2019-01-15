package gaea

import (
	"github.com/ttacon/libphonenumber"
	"strconv"
)

func GetPhoneNumberFromString(phone string) (*PhoneNumber, error) {
	num, err := libphonenumber.Parse(phone, "PE")
	if err != nil {
		return nil, err
	}
	n := int64(num.GetNationalNumber())
	number := strconv.FormatInt(n, 10)
	countryCode := strconv.Itoa(int(num.GetCountryCode()))
	return &PhoneNumber{
		CountryCode: countryCode,
		Number:      number,
		Valid:       false,
	}, nil
}
