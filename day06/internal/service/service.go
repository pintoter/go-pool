package service

import (
	"day06/internal/entity"
	"day06/internal/repository"

	"gorm.io/gorm"
)

type Repository interface {
	CreateArticle(title, content string) (int64, error)
	GetArticle(id int64) (entity.Article, error)
	GetArticles(id int64) ([]entity.Article, error)
}

type Service struct {
	repo Repository
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		repo: repository.NewRepo(db),
	}
}
