package http

import "errors"

var ErrInvalidToken = errors.New("authentication token is invalid or has expired")
