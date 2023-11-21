package entity

import "errors"

var (
	ErrArticleExists        = errors.New("article already exists")
	ErrArticleNotFound      = errors.New("article not found")
	ErrArticlesNotFound     = errors.New("articles not found")
	ErrEmptyParams          = errors.New("params can't be empty")
	ErrFailedCreateArticle  = errors.New("failed to create article")
	ErrFailedRenderTemplate = errors.New("failed render the template")
	ErrInvalidId            = errors.New("invalid id")
	ErrInvalidPage          = errors.New("invalid page")
	ErrInvalidCredentials   = errors.New("invalid credentials")
)
