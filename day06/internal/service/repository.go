package service

import (
	"context"
	"day06/internal/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func newRepo(db *gorm.DB) IRepository {
	return &repository{db: db}
}

func (r *repository) CreateArticle(ctx context.Context, name, link string) (int, error) {

	return 0, nil
}

func (r *repository) GetArticle(ctx context.Context, id int) (entity.Article, error) {

	return entity.Article{}, nil
}

func (r *repository) GetArticles(ctx context.Context, id int) ([]entity.Article, error) {

	return nil, nil
}
