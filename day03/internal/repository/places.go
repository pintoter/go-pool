package repository

import (
	"Day03/internal/entity"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

// var settings = `{
// 	"index" : {
// 			"max_result_window" : 20000
// 	}
// }`

type PlacesRepository struct {
	es *elasticsearch.Client
}

func NewPlacesRepository(es *elasticsearch.Client) *PlacesRepository {
	return &PlacesRepository{
		es: es,
	}
}

type searchResponseParams struct {
	Took    float64 `json:"took"`
	Timeout bool    `json:"timed_out"`
	Shards  struct {
		Total      int64 `json:"total"`
		Successful int64 `json:"successful"`
		Skipped    int64 `json:"skipped"`
		Failed     int64 `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int64  `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string        `json:"_index"`
			Id     string        `json:"_id"`
			Score  float64       `json:"_score"`
			Source *entity.Place `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (p *PlacesRepository) GetPlaces(limit int, offset int) ([]entity.Place, int, error) {
	resp, err := p.es.Search(
		p.es.Search.WithContext(context.Background()), // change ctx.Background() for ctx
		p.es.Search.WithIndex("places"),
		p.es.Search.WithFrom(offset),
		p.es.Search.WithSize(limit),
	)
	if err != nil {
		log.Println(err, "error in searching")
		return nil, 0, err
	}
	defer resp.Body.Close()

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		log.Println(err, "error in searching")
		return nil, 0, err
	}

	var respParams searchResponseParams
	err = json.NewDecoder(resp.Body).Decode(&respParams)
	if err != nil {
		log.Println(err, "error in decoding")
		return nil, 0, err
	}

	var places []entity.Place
	fmt.Println("Resp.Hit", respParams.Hits.Total.Value)
	if respParams.Hits.Total.Value > 0 {
		for _, hit := range respParams.Hits.Hits {
			if hit.Source == nil {
				log.Printf("hit with %s have nil Source", hit.Id)
				continue
			}
			places = append(places, *hit.Source)
		}
	}

	fmt.Println(places)

	return places, int(respParams.Hits.Total.Value), nil
}
