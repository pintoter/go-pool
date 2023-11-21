package service

import (
	"context"
	"day06/internal/entity"
	"log"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func newRepo(db *gorm.DB) IRepository {
	return &repository{db: db}
}

func (r *repository) CreateArticle(ctx context.Context, name, link string) (int, error) {
	dbArt := entity.Article{
		Title:   name,
		Content: link,
	}

	if err := r.db.Create(&dbArt).Error; err != nil {
		log.Println(err)
		return 0, err
	}

	return 0, nil
}

func (r *repository) GetArticle(ctx context.Context, id int) (entity.Article, error) {
	var article entity.Article

	if err := r.db.Where("id = ?", id).First(&article).Error; err != nil {
		return entity.Article{}, err
	}

	return article, nil
}

func (r *repository) GetArticles(ctx context.Context, limit, offset int) ([]entity.Article, int, error) {
	var total int64
	var articles []entity.Article

	if err := r.db.Order("id asc").Offset(offset).Limit(limit).Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	r.db.Model(&entity.Article{}).Count(&total)

	return articles, int(total), nil
}

func (r *repository) GetArticleByTitle(ctx context.Context, title string) (entity.Article, error) {
	var article entity.Article

	if err := r.db.Where("title = ?", title).First(&article).Error; err != nil {
		return entity.Article{}, err
	}

	return article, nil
}
