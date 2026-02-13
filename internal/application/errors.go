package application

import "errors"

var UserNotFoundError = errors.New("user not found")
var UserPasswordNotMatchError = errors.New("user password not match")
