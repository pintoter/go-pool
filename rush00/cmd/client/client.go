package main

import (
	"context"
	"errors"
	"fmt"
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
	client = "0.0.0.0:50055"
)

type Statistic struct {
	count int
	mean  float64
	sd    float64
}

func NewStat() *Statistic {
	return &Statistic{}
}

func (s *Statistic) Update(value float64) {
	s.count++

	delta := value - s.mean
	s.mean += delta / float64(s.count)
	s.UpdateSD()
}

func (s *Statistic) UpdateMean() {

}

func (s *Statistic) UpdateSD() {

}

type Message struct {
	ID        string
	Frequency float64
	Timestamp string
	Stat      *Statistic
}

func enterCoefficient(с chan<- float64) {
	var value float64
	fmt.Println("Enter coefficient:")
	fmt.Scanf("%f", &value)
	с <- value
}

func recieveMessages(wg *sync.WaitGroup, stream desc.Transmitter_TransmitClient) {
	var i int = 1

	// var dataPool = sync.Pool{
	// 	New: func() interface{} {
	// 		return &Message{
	// 			Stat: NewStat(),
	// 		}
	// 	},
	// }

	// Create client's signal for start
	var (
		coefficient  float64
		anomalyStage bool
	)
	anomalyStageSignal := make(chan os.Signal, 1)
	signal.Notify(anomalyStageSignal, syscall.SIGINT) // Handle client's signal "^ + \"

	chanWithFloat64 := make(chan float64)

	defer wg.Done()
	defer close(anomalyStageSignal)
	defer close(chanWithFloat64)

	for {
		select {
		case sig := <-anomalyStageSignal:
			log.Printf("got signal from client %v\n", sig)

			go enterCoefficient(chanWithFloat64)

			coefficient = <-chanWithFloat64
			anomalyStage = true
		default:
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

			// data := dataPool.Get().(*Message)

			// data.ID = resp.GetSessionId()
			// data.Frequency = resp.Frequency
			// data.Timestamp = resp.Timestamp

			// data.Stat.Update(resp.Frequency) // добавить обработку аномалий
			log.Printf("New recieve data #%d ID: %s\tFrequency: %f\tTimestamp: %s\n", i, resp.GetSessionId(), resp.GetFrequency(), resp.GetTimestamp())
			if anomalyStage {
				log.Println("MI V ANOMALII YRA")
				log.Println("KOEF =", coefficient)
			}
		}
	}
}

func main() {
	// Create context for handling client's signal with stop recieving data
	ctx, cancel := context.WithCancel(context.Background())
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGQUIT, syscall.SIGTERM) // ^ + \

	go func() {
		defer close(signals)
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
