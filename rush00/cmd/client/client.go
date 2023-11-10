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
	"time"

	desc "rush00/pkg/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	client = "0.0.0.0:50051"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-signals:
			log.Println("Recieve signal from client. Shutting down")
			cancel()
		}
	}()

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

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
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
			time.Sleep(2 * time.Second)
		}
	}()

	wg.Wait()

	if err := stream.CloseSend(); err != nil {
		log.Println(err)
	}

	log.Println("Client shutdown complete")
}
