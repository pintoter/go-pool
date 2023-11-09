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

type server struct {
	desc.UnimplementedGeneratorServer
}

func (s *server) GetData(ctx context.Context, req *desc.DataRequest) (*desc.DataResponse, error) {
	sessionID := uuid.New().String()
	frequency := rand.Float64()
	timestamp := time.Now().UTC().Format(time.RFC3339)

	resp := &desc.DataResponse{
		SessionId: sessionID,
		Frequency: frequency,
		Timestamp: timestamp,
	}

	fmt.Printf("Session ID: %s, Frequency: %f, Timestamp: %s\n", sessionID, frequency, timestamp)

	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterGeneratorServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
