package service

import (
	"rush00/internal/statistic"
	desc "rush00/pkg/api/proto"

	"gorm.io/gorm"
)

type Reciever interface {
	Recieve(stream desc.Transmitter_TransmitClient, coefficient float64)
}

type Service struct {
	Reciever
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		Reciever: NewRecieverService(db, statistic.NewStat()),
	}
}
