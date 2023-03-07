package util

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid Password")
	ErrUserAlreadyExists = errors.New("such user already registered")
)
