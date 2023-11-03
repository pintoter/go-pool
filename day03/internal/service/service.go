package service

import (
	"Day03/internal/entity"
)

type Store interface {
	GetPlaces(limit int, offset int) ([]entity.Place, int, error)
}

type Service struct {
	Store
}

func NewService() *Service {
	return &Service{
		Store: NewStoreService(),
	}
}
