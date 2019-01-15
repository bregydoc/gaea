package gaea

import (
	"github.com/ttacon/libphonenumber"
	"strconv"
	"strings"
)

func getPhoneNumberFromString(phone string) (*PhoneNumber, error) {
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

func capitalizeString(value string) string {
	if len(value) < 1 {
		return ""
	}

	if len(value) == 1 {
		return strings.ToUpper(value)
	}

	head := strings.ToUpper(string(value[0]))
	tail := value[1:]
	return head + tail
}
