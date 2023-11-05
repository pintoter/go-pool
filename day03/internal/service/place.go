package service

import (
	"Day03/internal/entity"
	"Day03/internal/repository"
)

type PlacesService struct {
	placeRepo repository.Places
}

func NewPlacesService(repo repository.Places) *PlacesService {
	return &PlacesService{
		placeRepo: repo,
	}
}

func (s *PlacesService) GetPlaces(limit int, offset int) ([]entity.Place, int, error) {
	places, total, err := s.placeRepo.GetPlaces(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return places, total, nil
}

func (s *PlacesService) GetClosestPlaces(lat, lon float64) ([]entity.Place, error) {
	places,err := s.placeRepo.GetClosestPlace(lat, lon)
	if err != nil {
		return nil, err
	}

	return places, nil
}
