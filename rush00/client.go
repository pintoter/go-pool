package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	desc "rush00/pkg/api/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := dest.NewGeneratorClient(conn)
	resp, err := client.GetData(context.Background(), &dest.DataRequest{})
	if err != nil {
		log.Println(err)
	}
	
	fmt.Printf("ID: %s\nFrequency:%f\nTimestamp:%s", resp.GetSessionId(), resp.GetFrequency(), resp.GetTimestamp())
}
