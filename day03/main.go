package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

var (
	ESUsername    = "elastic"
	ESPW          = "Wqy0O9nLa8HCHZ*7MHLA"
	ESFingerPrint = "7140373329f32249e4775f12a8d4e35f5f2b2c390f4677aaf4f4b7c48c469ee9"
	address       = "https://localhost:9200"
)

var (
	file = "../materials/data1.csv"
)

const (
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

type Place struct {
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	Location GeoPoint `json:"location"`
}

type GeoPoint struct {
	Longitude string `json:"lon"`
	Latitude  string `json:"lat"`
}

func main() {
	ctx := context.Background()

	// Create ES client
	es, err := ConnectWithES()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to ES")

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
			Action:     index,
			DocumentID: line[0],
			Body:       bytes.NewReader(record),
			OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, respb esutil.BulkIndexerResponseItem) {
				log.Println("successfully added new item:", item, respb)
			},
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, respb esutil.BulkIndexerResponseItem, err error) {
				log.Fatal("failed to add new item:", item, respb, err)
			},
		})
		if err != nil {
			log.Fatal(err)
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

	for i := 0; ; i++ {
		_, err := es.Get("places", strconv.Itoa(i))
		if err != nil {
			log.Println("error getting document from database:", err)
			break
		}
	}
}

func newRecord(line []string) ([]byte, error) {
	var place Place
	place.Name = line[1]
	place.Address = line[2]
	place.Phone = line[3]
	place.Location = GeoPoint{
		Longitude: line[4],
		Latitude:  line[5],
	}
	res, err := json.Marshal(place)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ConnectWithES() (*elasticsearch.Client, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			address,
		},
		Username:               ESUsername,
		Password:               ESPW,
		CertificateFingerprint: ESFingerPrint,
		Logger: &estransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	})
	if err != nil {
		log.Fatalln("error with connecting to ES:", err)
		return nil, err
	}

	return client, nil
}
