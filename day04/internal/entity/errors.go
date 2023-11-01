package entity

import "errors"

var (
	NotEnoughMoney   = errors.New("You need %d more money!")
	ErrorInInputData = errors.New("Some error in input data")
)
