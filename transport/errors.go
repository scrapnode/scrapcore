package transport

import "errors"

var (
	ErrBodyInvalid = errors.New("transport: could not parser request body")
)
