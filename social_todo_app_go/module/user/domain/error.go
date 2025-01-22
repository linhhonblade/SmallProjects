package domain

import "errors"

var (
	ErrEmailHasExisted   = errors.New("email has been existed")
	InvalidEmailPassword = errors.New("invalid email and password")
)
