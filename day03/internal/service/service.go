package service

import (
	"Day03/internal/entity"
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

const (
	file    = "../materials/data1.csv"
	index   = "places"
	mapping = `{
	"mappings": {
		"properties": {
			"name": {
				"type": "text"
			},
			"address": {
				"type": "text"
			},
			"phone": {
				"type": "text"
			},
			"location": {
				"type": "geo_point"
			}
		}
	}
}`
)

func LoadData(ctx context.Context, es *elasticsearch.Client) {
	resp, err := esapi.IndicesExistsRequest{
		Index: []string{index},
	}.Do(ctx, es)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		// Delete for index exist
		_, err = esapi.IndicesDeleteRequest{
			Index: []string{index},
		}.Do(ctx, es)
		if err != nil {
			log.Println(err)
		}
		log.Println("Index deleted or not exists")
	}

	// Create Request
	resp, err = esapi.IndicesCreateRequest{
		Index:  index,
		Body:   strings.NewReader(mapping),
		Pretty: true,
		Human:  true,
	}.Do(ctx, es)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	log.Println("Created request to ES")

	// Read data from data.csv
	fi, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	reader := csv.NewReader(fi)
	reader.Comma = '\t'

	// Create Bulk
	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  index,
		Client: es,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer bulkIndexer.Close(ctx)

	for {
		line, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		record, err := newRecord(line)
		if err != nil {
			log.Println(err)
			break
		}

		err = bulkIndexer.Add(ctx, esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: line[0],
			Body:       bytes.NewReader(record),
			OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, respb esutil.BulkIndexerResponseItem) {
				log.Println("successfully added new item:", item, respb)
			},
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, respb esutil.BulkIndexerResponseItem, err error) {
				log.Println("failed to add new item:", item, respb, err)
			},
		})
		if err != nil {
			log.Println(err)
		}
	}
	err = bulkIndexer.Close(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if bulkIndexer.Stats().NumFailed > 0 {
		log.Fatalln(bulkIndexer.Stats().NumFailed)
	}
	log.Println(bulkIndexer.Stats())

}

func newRecord(line []string) ([]byte, error) {
	var place entity.Place
	place.Name = line[1]
	place.Address = line[2]
	place.Phone = line[3]
	place.Location = entity.GeoPoint{
		Longitude: line[4],
		Latitude:  line[5],
	}
	res, err := json.Marshal(place)
	if err != nil {
		return nil, err
	}
	return res, nil
}

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
