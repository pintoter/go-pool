package repository

import (
	"Day03/internal/entity"

	"github.com/elastic/go-elasticsearch/v7"
)

type Places interface {
	GetPlaces(limit int, offset int) ([]entity.Place, int, error)
}

type Repository struct {
	Places
}

func NewRepository(es *elasticsearch.Client) *Repository {
	return &Repository{
		Places: NewPlacesRepository(es),
	}
}
