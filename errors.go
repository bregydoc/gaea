package gaea

import "errors"

// ErrAccountTypeNotExist is an error type of gaea
var ErrAccountTypeNotExist = errors.New("account type not defined or not exist")

// ErrInvalidPhoneNumber is an error type of gaea
var ErrInvalidPhoneNumber = errors.New("invalid phone number")

// ErrUnimplementedError is an error type of gaea
var ErrUnimplementedError = errors.New("unimplemented, TODO: implement please")

// ErrInvalidPersonName is an error type of gaea
var ErrInvalidPersonName = errors.New("user name are not valid, please set the name like \"Jon Doe\"")

// ErrInvalidSignUpPerson is an error type of gaea
var ErrInvalidSignUpPerson = errors.New("invalid sign up, person have more than one account and we need a NEW person")

// ErrAccountPersonNoMatch is an error type of gaea
var ErrAccountPersonNoMatch = errors.New("invalid account, owner is not the same person")
