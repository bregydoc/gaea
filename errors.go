package gaea

import "errors"

var accountTypeNotExist = errors.New("account type not defined or not exist")
var invalidPhoneNumber = errors.New("invalid phone number")
var unimplementedError = errors.New("unimplemented, TODO: implement please")
var invalidPersonName = errors.New("user name are not valid, please set the name like \"Jon Doe\"")
