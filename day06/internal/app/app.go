package app

import (
	"day06/internal/config"
	"day06/internal/server"
	"day06/internal/service"
	"day06/internal/transport"
	"day06/pkg/db/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	configPath = "./configs/admin_credentials.txt"
)

func Run() {
	err := config.ReadConfigTxt(configPath)
	if err != nil {
		log.Fatal("Failed init configuration:", err)
	}
	cfg := config.GetConfigInstance()

	db, err := postgres.ConnectDB(&cfg.DB)
	if err != nil {
		log.Fatal("Failed connect to database:", err)
	}

	service := service.NewService(db)

	handler := transport.NewHandler(service)

	server := server.New(handler, cfg.Addr.Host, cfg.Addr.Port)

	server.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case s := <-quit:
		log.Printf("Starting gracefully shutdown after signal %s", s.String())
	case err = <-server.Notify():
		log.Fatalf("Error when starting server: %s", err.Error())
	}

	if err := server.Shutdown(); err != nil {
		log.Fatal("Server", err)
	}

	log.Println("Server gracefully shutting down")
}
