package entity

import "errors"

var (
	InvalidQuery = errors.New("Invalid 'page' value: %s")
)
