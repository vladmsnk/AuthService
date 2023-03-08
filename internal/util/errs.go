package util

import "errors"

var (
	ErrUserNotFound               = errors.New("user not found")
	ErrInvalidPassword            = errors.New("invalid Password")
	ErrUserAlreadyExists          = errors.New("such user already registered")
	ErrInvalidEmailFormat         = errors.New("email format is not valid")
	ErrUnexpectedSigningAlgorithm = errors.New("unexpected signing method")
	ErrInvalidAccessToken         = errors.New("invalid access token")
	ErrExpiredToken               = errors.New("token is expired")
)
