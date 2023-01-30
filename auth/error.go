package auth

import "errors"

var (
	ErrSignFailed         = errors.New("auth: incorrect username or password")
	ErrInvalidSignMethod  = errors.New("auth: unexpected signing method")
	ErrInvalidToken       = errors.New("auth: token is not valid")
	ErrInvalidTokenClaims = errors.New("auth: token claims is not valid")
	ErrInvalidTokenType   = errors.New("auth: token type is not valid")
)
