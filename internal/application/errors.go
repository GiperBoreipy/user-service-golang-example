package application

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrUserPasswordNotMatch = errors.New("user password not match")
