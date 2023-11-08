package app

import (
	"context"
	"day04/internal/server"
	"day04/internal/service"
	"day04/internal/transport"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	host = "127.0.0.1"
	port = "3333"
)

func Run(isSecure bool) {
	services := service.New()

	handler := transport.NewHandler(services)

	server := server.New(handler, host, port)

	go func() {
		var err error

		if !isSecure {
			err = server.Run()
			log.Println("Starting HTTP server")
		} else if isSecure {
			err = server.RunTLS()
			log.Println("Starting HTTPS server")
		}

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

	err := server.Stop(ctx)
	if err != nil {
		log.Printf("Server shutdown failed:%v", err)
	}
	log.Println("Graceful shutdown completed")
}
