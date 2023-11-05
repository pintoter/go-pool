package app

import (
	"Day03/internal/repository"
	"Day03/internal/server"
	"Day03/internal/service"
	"Day03/internal/transport"
	database "Day03/pkg/database/elasticsearch"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	ESUsername    = "elastic"
	ESPW          = "Wqy0O9nLa8HCHZ*7MHLA"
	ESFingerPrint = "7140373329f32249e4775f12a8d4e35f5f2b2c390f4677aaf4f4b7c48c469ee9"
	address       = "http://localhost:9200"
	host          = "127.0.0.1"
	port          = "8888"
)

func Run() {
	es, err := database.New(address)
	if err != nil {
		log.Fatal(err)
	}

	services := service.New(service.Deps{
		Repos: repository.New(es),
	})

	handler := transport.NewHandler(services)

	server := server.New(handler, host, port)
	go func() {
		err := server.Run()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server ListenAndServe: %v", err)
		}
		log.Println("Server successfully stopped")
	}()

	quitChan := make(chan os.Signal)
	signal.Notify(quitChan, syscall.SIGTERM, syscall.SIGINT)
	quitSig := <-quitChan

	log.Println("Graceful shutdown starter with signal:", quitSig)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = server.Stop(ctx)
	if err != nil {
		log.Printf("Server shutdown failed:%v", err)
	}
	log.Println("Graceful shutdown completed")
}
