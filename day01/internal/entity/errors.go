package entity

import "errors"

var (
	EmptyFilename    = errors.New("entry the filename")
	FileNotFound     = errors.New("file doesn't exists")
	UnexpectedFormat = errors.New("unknown format")
)
