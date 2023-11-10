package main

import (
	"context"
	"errors"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	desc "rush00/pkg/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	client = "0.0.0.0:50051"
)

type data struct {
	ID        string
	Frequency float64
	Timestamp string
}

func recieveMessages(wg *sync.WaitGroup, stream desc.Transmitter_TransmitClient) {
	var i int = 1
	defer wg.Done()
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			log.Println(err)
			return
		}

		if status.Code(err) == codes.Canceled {
			log.Println("Server's stream closed after signal from client")
			return
		}

		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("New recieve data #%d ID: %s\tFrequency: %f\tTimestamp: %s\n", i, resp.GetSessionId(), resp.GetFrequency(), resp.GetTimestamp())
		i++
	}
}

func main() {
	// Create context for handling client's signal with stop recieving data
	ctx, cancel := context.WithCancel(context.Background())
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-signals:
			log.Println("Recieve signal from client. Shutting down")
			cancel()
		}
	}()

	// Create gRPC connect
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(client, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := desc.NewTransmitterClient(conn)
	stream, err := client.Transmit(ctx, &desc.DataRequest{})
	if err != nil {
		log.Println(err)
	}
	log.Println("gRPC client started")

	var wg sync.WaitGroup

	wg.Add(1)
	go recieveMessages(&wg, stream)
	wg.Wait()

	log.Println("Client shutdown complete")
}
