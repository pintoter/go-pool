package database

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
	file    = "./config/data.csv"
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
	settings = `{
	"index.max_result_window" : 20000
}`
)

var (
	ESUsername    = "elastic"
	ESPW          = "Wqy0O9nLa8HCHZ*7MHLA"
	ESFingerPrint = "7140373329f32249e4775f12a8d4e35f5f2b2c390f4677aaf4f4b7c48c469ee9"
)

func New(address string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
		Username:               ESUsername,
		Password:               ESPW,
		CertificateFingerprint: ESFingerPrint,
		// Logger: &estransport.ColorLogger{
		// 	Output:             os.Stdout,
		// 	EnableRequestBody:  true,
		// 	EnableResponseBody: true,
		// },
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// Check ES for existing index

	// resp, err := esapi.IndicesPutSettingsRequest{
	// 	Index: []string{"places"},
	// 	Body:  strings.NewReader(settings),
	// }.Do(context.Background(), es)
	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Body.Close()

	err = UpdateMaxResultSettings(es)
	LoadData(es)
	if err != nil {
		return nil, err
	}

	return es, nil
}

func UpdateMaxResultSettings(es *elasticsearch.Client) error {
	resp, err := esapi.IndicesPutSettingsRequest{
		Index: []string{"places"},
		Body:  strings.NewReader(settings),
	}.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func LoadData(es *elasticsearch.Client) {
	ctx := context.Background()

	resp, err := esapi.IndicesExistsRequest{
		Index: []string{index},
	}.Do(ctx, es)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		// Delete for index exist
		return
		// _, err = esapi.IndicesDeleteRequest{
		// 	Index: []string{index},
		// }.Do(ctx, es)
		// if err != nil {
		// 	log.Println(err)
		// }
		// log.Println("Index deleted or not exists")
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
		})
		if err != nil {
			log.Println(err)
		}
	}
	if bulkIndexer.Stats().NumFailed > 0 {
		log.Println(bulkIndexer.Stats().NumFailed)
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
