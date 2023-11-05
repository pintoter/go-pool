package repository

import (
	"Day03/internal/entity"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

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

type searchRequestParams struct {
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
	// 	settings := `{
	//     "size" : 20000
	// }`

	// 	resp, err := esapi.SearchRequest{
	// 		Index: []string{"places"},
	// 		Body:  strings.NewReader(settings),
	// 	}.Do(context.Background(), p.es)
	// 	if err != nil {
	// 		return nil, 0, err
	// 	}
	// 	defer resp.Body.Close()

	resp, err := p.es.Search(
		p.es.Search.WithContext(context.Background()),
		p.es.Search.WithIndex("places"),
		p.es.Search.WithFrom(offset),
		p.es.Search.WithSize(limit),
	)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	var respParams searchRequestParams
	err = json.NewDecoder(resp.Body).Decode(&respParams)
	if err != nil {
		return nil, 0, err
	}

	var places []entity.Place
	// log.Println("Resp.Hit", respParams.Hits.Total.Value)
	if respParams.Hits.Total.Value > 0 {
		for _, hit := range respParams.Hits.Hits {
			if hit.Source == nil {
				log.Printf("hit with %s have nil Source", hit.Id)
				continue
			}
			places = append(places, *hit.Source)
		}
	}

	return places, int(respParams.Hits.Total.Value), nil
}

func (p *PlacesRepository) GetClosestPlace(lat, lon float64) ([]entity.Place, error) {
	query := `{
		"from": ` + fmt.Sprintf("%d", 0) + `,
		"size": ` + fmt.Sprintf("%d", 3) + `,
		"sort": [
				{
						"_geo_distance": {
								"location": {
										"lat": ` + fmt.Sprintf("%f", lat) + `,
										"lon": ` + fmt.Sprintf("%f", lon) + `
								},
								"order": "asc",
								"unit": "km",
								"mode": "min",
								"distance_type": "arc",
								"ignore_unmapped": true
						}
				}
		]
}`
	res, err := p.es.Search(
		p.es.Search.WithContext(context.Background()),
		p.es.Search.WithIndex("places"),
		p.es.Search.WithBody(strings.NewReader(query)),
		p.es.Search.WithSize(3),
	)
	if err != nil {
		log.Printf("Error executing the search request: %s", err)
	}

	var respParams searchRequestParams
	err = json.NewDecoder(res.Body).Decode(&respParams)
	if err != nil {
		return nil, err
	}

	var places []entity.Place
	// log.Println("Resp.Hit", respParams.Hits.Total.Value)
	if respParams.Hits.Total.Value > 0 {
		for _, hit := range respParams.Hits.Hits {
			if hit.Source == nil {
				log.Printf("hit with %s have nil Source", hit.Id)
				continue
			}

			var place entity.Place
			place.ID = hit.Id
			place.Name = hit.Source.Name
			place.Address = hit.Source.Address
			place.Phone = hit.Source.Phone
			place.Location.Latitude = hit.Source.Location.Latitude
			place.Location.Longitude = hit.Source.Location.Longitude

			places = append(places, place)
		}
	}

	return places, nil
}
