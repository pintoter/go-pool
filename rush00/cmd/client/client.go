package main

import (
	"context"
	"errors"
	"io"
	"log"

	desc "rush00/pkg/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	client = "0.0.0.0:50051"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(client, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := desc.NewTransmitterClient(conn)
	stream, err := client.Transmit(context.Background(), &desc.DataRequest{})
	if err != nil {
		log.Println(err)
	}

	done := make(chan bool)

	go func() {
		var i int = 1
		for {
			resp, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				done <- true
				return
			}

			if err != nil {
				log.Fatalln("can't recieve data:", err)
			}

			log.Printf("New recieve data #%d ID: %s\tFrequency: %f\tTimestamp: %s\n", i, resp.GetSessionId(), resp.GetFrequency(), resp.GetTimestamp())
			i++
		}
	}()

	<-done

	log.Println("stream has ended")
}
