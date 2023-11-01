package repository

import (
	"Day03/internal/entity"

	"github.com/elastic/go-elasticsearch/v7"
)

type PlacesRepository struct {
	es *elasticsearch.Client
}

func NewPlacesRepository(es *elasticsearch.Client) *PlacesRepository {
	return &PlacesRepository{
		es: es,
	}
}

func (p *PlacesRepository) GetPlaces(limit int, offset int) ([]entity.Place, int, error) {

	return nil, 0, nil
}
