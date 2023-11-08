package app

import (
	"context"
	"day04/internal/config"
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

func Run(cfg *config.Config, isSecure bool) {
	services := service.New()

	handler := transport.NewHandler(services)

	server := server.New(handler, cfg.Host, cfg.Port)

	go func() {
		var err error

		if !isSecure {
			err = server.Run()
			log.Println("Starting HTTP server")
		} else {
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
