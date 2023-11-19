package service

import (
	"context"
	"day06/internal/entity"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

type IRepository interface {
	CreateArticle(ctx context.Context, name, link string) (int, error)
	GetArticle(ctx context.Context, id int) (entity.Article, error)
	GetArticles(ctx context.Context, id int) ([]entity.Article, error)
}

type Service struct {
	repo IRepository
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		repo: newRepo(db),
	}
}

func (s *Service) CreateArticle(ctx context.Context, name, link string) (int, error) {
	isExists, err := s.isArticleExists(ctx, link)
	if err != nil {
		log.Println(err)
	}

	if isExists {
		return 0, errors.New("article already exists")
	}

	mdFile := fmt.Sprintf("./ui/md/%s", name)
	if _, err := os.Create(mdFile); err != nil {
		log.Println(err)
		return 0, err
	}

	htmlName := strings.Split(name, ".md")[0] + ".html"
	htmlFile := fmt.Sprintf("./ui/html/%s", htmlName)
	if _, err := os.Create(htmlFile); err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := s.repo.CreateArticle(ctx, name, link)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (s *Service) GetArticle(ctx context.Context) (entity.Article, error) {

	return entity.Article{}, nil
}

func (s *Service) GetArticles(ctx context.Context) ([]entity.Article, error) {

	return nil, nil
}

func (s *Service) isArticleExists(ctx context.Context, link string) (bool, error) {

	return true, nil
}
