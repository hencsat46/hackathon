package exceptions

import "errors"

var (
	ErrInternalServerError  = errors.New("internal server error")
	ErrBadRequest           = errors.New("bad request")
	ErrPasswordIncorrect    = errors.New("password incorrect")
	ErrNotFound             = errors.New("user not found")
	ErrInvalidSigningMethod = errors.New("provided token has been signed with invalid method")
	ErrInvalidToken         = errors.New("provided token is invalid")
)
