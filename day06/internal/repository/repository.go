package repository

import (
	"day06/internal/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateArticle(title, content string) (int64, error) {

	return 0, nil
}

func (r *repository) GetArticle(id int64) (entity.Article, error) {

	return entity.Article{}, nil
}

func (r *repository) GetArticles(id int64) ([]entity.Article, error) {

	return nil, nil
}
