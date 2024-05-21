package exceptions

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrBadRequest          = errors.New("bad request")
	ErrPasswordIncorrect   = errors.New("password incorrect")
	ErrNotFound            = errors.New("user not found")
)
