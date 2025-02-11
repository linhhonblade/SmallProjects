package domain

import "errors"

var (
	ErrEmailHasExisted    = errors.New("email has been existed")
	InvalidEmailPassword  = errors.New("invalid email and password")
	ErrCannotChangeAvatar = errors.New("cannot change your avatar")
)
