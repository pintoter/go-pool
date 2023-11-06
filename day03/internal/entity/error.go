package entity

import "errors"

var (
	InvalidQuery = errors.New("invalid 'page' value: %s")
	UnauthorizedToken = errors.New("unauthorized or token expired, get a new token")
)
