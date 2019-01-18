package gaea

import "errors"

var errAccountTypeNotExist = errors.New("account type not defined or not exist")
var errInvalidPhoneNumber = errors.New("invalid phone number")
var errUnimplementedError = errors.New("unimplemented, TODO: implement please")
var errInvalidPersonName = errors.New("user name are not valid, please set the name like \"Jon Doe\"")

var errInvalidSignUpPerson = errors.New("invalid sign up, person have more than one account and we need a NEW person")
var errAccountPersonNoMatch = errors.New("invalid account, owner is not the same person")
