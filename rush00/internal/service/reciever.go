package service

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"rush00/internal/statistic"
	"strconv"
	"sync"
	"syscall"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	desc "rush00/pkg/api/proto"
	"rush00/pkg/database/postgres"
)

type RecieverService struct {
	db   *gorm.DB
	stat *statistic.Statistic
}

func NewRecieverService(db *gorm.DB, stat *statistic.Statistic) *RecieverService {
	return &RecieverService{
		db:   db,
		stat: stat,
	}
}

func (r *RecieverService) Recieve(stream desc.Transmitter_TransmitClient, coefficient float64) {
	var (
		anomalyStage bool
		pool         = &sync.Pool{
			New: func() interface{} {
				return &desc.DataResponse{}
			},
		}
	)

	anomalyStageSignal := make(chan os.Signal, 1)
	signal.Notify(anomalyStageSignal, syscall.SIGINT)

	for {
		select {
		case sig := <-anomalyStageSignal:
			log.Printf("got signal from client %v\n", sig)

			if coefficient == 0 {
				chanWithFloat64 := make(chan float64)
				defer close(chanWithFloat64)
				go func() {
					var value float64
					fmt.Println("Enter coefficient:")
					fmt.Scanf("%f", &value)
					chanWithFloat64 <- value
				}()
				coefficient = <-chanWithFloat64
			}
			anomalyStage = true
			log.Println("starting anomaly stage")
		default:
			var err error
			resp := pool.Get().(*desc.DataResponse)
			resp, err = stream.Recv()
			if errors.Is(err, io.EOF) {
				log.Println(err)
				return
			}
			if status.Code(err) == codes.Canceled {
				log.Println("Server's stream closed after signal from client")
				return
			}
			if status.Code(err) == codes.Internal {
				log.Println("Server unavailible. Closing client...")
				return
			}
			if err != nil {
				log.Println(err)
				return
			}

			log.Printf(
				"New data recievedID: %s\tFrequency: %f\tTimestamp: %s\n",
				resp.GetSessionId(),
				resp.GetFrequency(),
				resp.GetTimestamp(),
			)
			if anomalyStage {
				if r.isAnomaly(resp.Frequency, coefficient) {
					timestmp, _ := strconv.Atoi(resp.GetTimestamp())
					err = r.db.Create(&postgres.Message{
						SessionID: resp.GetSessionId(),
						Frequency: resp.GetFrequency(),
						Timestamp: uint64(timestmp),
					}).Error
					if err != nil {
						log.Println("Error writing to database:", err)
						return
					}
				}
				continue
			} else {
				r.stat.Update(resp.Frequency)
				log.Printf("New data successfully processed: Count: %d\tMean: %f\tStandart Deaviation: %f\n", r.stat.Count, r.stat.Mean, r.stat.SD)
			}
			resp.Reset()
			pool.Put(resp)
		}
	}
}

func (r *RecieverService) isAnomaly(value, coefficient float64) bool {
	low := r.stat.Mean - (coefficient * r.stat.SD)
	high := r.stat.Mean + (coefficient * r.stat.SD)

	if !(value >= low && value <= high) {
		log.Printf("New anomaly detected: Value: %f\t Low: %f High: %f\n", value, low, high)
		return true
	}

	return false
}
