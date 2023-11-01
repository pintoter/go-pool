package app

import (
	"Day03/internal/service"
	"Day03/pkg/database"
	"context"
	"log"
)

var (
	ESUsername    = "elastic"
	ESPW          = "Wqy0O9nLa8HCHZ*7MHLA"
	ESFingerPrint = "7140373329f32249e4775f12a8d4e35f5f2b2c390f4677aaf4f4b7c48c469ee9"
	address       = "http://localhost:9200"
)

func Run() {
	ctx := context.Background()

	es, err := database.NewDB(address)
	if err != nil {
		log.Fatal(err)
	}

	service.LoadData(ctx, es)
}
