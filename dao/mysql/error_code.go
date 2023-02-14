package mysql

import "errors"

var (
	ErrorUserExist                   = errors.New("user existed")
	ErrorIncorrectUsernameOrPassword = errors.New("Incorrect username or password")
	ErrorInvalidID                   = errors.New("Invalid ID")
)
