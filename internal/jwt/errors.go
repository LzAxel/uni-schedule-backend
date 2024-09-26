package jwt

import (
	"errors"
)

var (
	ErrTokenExpired = errors.New("token is expired")
	ErrInvalidToken = errors.New("invalid token")
)
