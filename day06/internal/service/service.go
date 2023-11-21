package service

import (
	"context"
	"day06/internal/entity"
	"errors"
	"log"

	"gorm.io/gorm"
)

type IRepository interface {
	CreateArticle(ctx context.Context, name, link string) (int, error)
	GetArticle(ctx context.Context, id int) (entity.Article, error)
	GetArticles(ctx context.Context, limit, offset int) ([]entity.Article, int, error)
	GetArticleByTitle(ctx context.Context, title string) (entity.Article, error)
}

type Service struct {
	repo IRepository
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		repo: newRepo(db),
	}
}

func (s *Service) CreateArticle(ctx context.Context, title, content string) (int, error) {
	isExists := s.isArticleExists(ctx, title)

	if isExists {
		return 0, errors.New("article already exists")
	}

	id, err := s.repo.CreateArticle(ctx, title, content)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (s *Service) GetArticle(ctx context.Context, id int) (entity.Article, error) {
	return s.repo.GetArticle(ctx, id)
}

func (s *Service) GetArticles(ctx context.Context, limit, offset int) ([]entity.Article, int, error) {
	return s.repo.GetArticles(ctx, limit, offset)
}

func (s *Service) isArticleExists(ctx context.Context, title string) bool {
	_, err := s.repo.GetArticleByTitle(ctx, title)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
