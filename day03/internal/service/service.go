package service

import (
	"Day03/internal/entity"
	"Day03/internal/repository"
)

type Store interface {
	GetPlaces(limit int, offset int) ([]entity.Place, int, error)
	GetClosestPlaces(lat, lon float64) ([]entity.Place, error)
}

type Service struct {
	Store
}

type Deps struct {
	Repos *repository.Repository
}

func New(deps Deps) *Service {
	return &Service{
		Store: NewPlacesService(deps.Repos.Places),
	}
}
