package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	desc "rush00/pkg/api/proto"
	"rush00/pkg/database/postgres"

	"rush00/internal/config"
	"rush00/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddr = "0.0.0.0:50055"
)

var coefficient float64

type grpcClient struct {
	service           *service.Service
	transmitterClient desc.TransmitterClient
}

func NewGrpcClient(service *service.Service, conn *grpc.ClientConn) *grpcClient {
	return &grpcClient{
		service:           service,
		transmitterClient: desc.NewTransmitterClient(conn),
	}
}

func init() {
	flag.Float64Var(&coefficient, "k", 0, "Anomaly STD coefficient")
	flag.Parse()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		defer close(signals)
		select {
		case <-signals:
			log.Println("Recieve signal from client. Shutting down")
			cancel()
		}
	}()

	cfg := config.Get()

	db, err := postgres.New(&cfg.DB)
	if err != nil {
		log.Fatal("failed connecting to DB", err)
	}

	service := service.NewService(db)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(cfg.Client.Address, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewGrpcClient(service, conn)
	stream, err := client.transmitterClient.Transmit(ctx, &desc.DataRequest{})
	if err != nil {
		log.Println(err)
	}
	log.Println("gRPC client started")

	service.Recieve(stream, coefficient)

	log.Println("Client shutdown complete")
}
