package errors

import "errors"

var UserEmailNotValidError = errors.New("user email not valid")
var UserNotFoundError = errors.New("user not found")
var UserPasswordNotMatchError = errors.New("user password not match")
