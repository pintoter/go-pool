package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	desc "rush00/pkg/api/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	meanMin      float64 = -10.0
	meanMax      float64 = 10
	sdMin        float64 = 0.3
	sdMax        float64 = 1.5
	grpcHostPort         = "0.0.0.0:50051"
)

type GrpcServer struct {
	desc.UnimplementedTransmitterServer
}

func (s *GrpcServer) Transmit(req *desc.DataRequest, stream desc.Transmitter_TransmitServer) error {
	// Set uuID
	sessionID := uuid.New().String()

	// Generate const mean and sd for request
	mean := meanMin + rand.Float64()*meanMax
	sd := sdMin + rand.Float64()*(sdMax-sdMin)
	randomGen := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	log.Println("Generating data for new request")

	var i int = 1

	for {
		select {
		case <-stream.Context().Done():
			log.Println("stream has ended")
			return status.Error(codes.Canceled, "stream has ended")
		default:
			frequency := randomGen.NormFloat64()*sd + mean

			message := &desc.DataResponse{
				SessionId: sessionID,
				Frequency: frequency,
				Timestamp: time.Now().Format(time.RFC3339),
			}

			log.Printf("Generation #%d:\tID: %s\tFrequency: %f\tTimestamp: %s\n", i, message.SessionId, message.Frequency, message.Timestamp)

			if err := stream.SendMsg(message); err != nil {
				return status.Error(codes.Canceled, err.Error())
			}

			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", grpcHostPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	reflection.Register(grpcServer)
	desc.RegisterTransmitterServer(grpcServer, &GrpcServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// add graceful shutdown
}
