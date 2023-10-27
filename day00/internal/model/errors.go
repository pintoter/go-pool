package model

import "errors"

var (
	DigitOverLimit = errors.New("Number out of bouds")
	NotDigitInput  = errors.New("Input must be integer")
	EmptyInput     = errors.New("Empty input")
)
