package main

import (
	"log"
	"net"

	desc "rush00/pkg/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	desc.UnimplementedGeneratorServer
}

func (s *server) GetGenerateData(req *desc.DataRequest, nn desc.Generator_GetGenerateDataServer) error { // почему так сгенерировалось?

	return nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:")
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
