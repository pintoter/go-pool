package service

import "Day03/internal/entity"

type StoreService struct{}

func NewStoreService() *StoreService {
	return &StoreService{}
}

func (s *StoreService) GetPlaces(limit int, offset int) ([]entity.Place, int, error) {

	return nil, 0, nil
}
