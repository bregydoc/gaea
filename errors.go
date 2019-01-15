package gaea

import "errors"

var accountTypeNotExist = errors.New("account type not defined or not exist")
var invalidPhoneNumber = errors.New("invalid phone number")
var unimplementedError = errors.New("unimplemented, TODO: implement please")
