package entity

import "errors"

var (
	ArgumentIsEmpty    = errors.New("path is empty")
	UncorrectExt       = errors.New("-ext must be used only with flag -f")
	TooMuchFlags       = errors.New("you can use just 1 flag")
	PathIsNotDirectory = errors.New("path isn't directory")
	EmptyFiles         = errors.New("The list of files is empty")
)
