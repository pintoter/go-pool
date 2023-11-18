package main

import (
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
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
	grpcHostPort         = "0.0.0.0:50055"
)

var stopServer = make(chan struct{}, 1)

type GrpcServer struct {
	desc.UnimplementedTransmitterServer
}

func (s *GrpcServer) Transmit(req *desc.DataRequest, stream desc.Transmitter_TransmitServer) error {
	sessionID := uuid.New().String()

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
		case <-stopServer:
			go func() {
				time.Sleep(3 * time.Second)
				stream.Send(nil)
			}()
		default:
			frequency := randomGen.NormFloat64()*sd + mean

			message := &desc.DataResponse{
				SessionId: sessionID,
				Frequency: frequency,
				Timestamp: strconv.Itoa(int(time.Now().Unix())),
			}

			log.Printf(
				"Generation #%d:\tID: %s\tFrequency: %f\tTimestamp: %s\n",
				i,
				message.SessionId,
				message.Frequency,
				message.Timestamp,
			)

			if err := stream.SendMsg(message); err != nil {
				return err
			}

			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", grpcHostPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	desc.RegisterTransmitterServer(grpcServer, &GrpcServer{})

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		sig := <-quit
		log.Printf("got signal %v, attempting graceful shutdown", sig)
		stopServer <- struct{}{}
		grpcServer.Stop()
	}()

	log.Println("starting gRPC server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	wg.Wait()
	log.Println("gRPC server's stop")
}
