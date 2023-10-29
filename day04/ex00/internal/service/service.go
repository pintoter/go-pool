package service

type Candies interface {
	Buy(name string, countCandy, money int64) (int64, error)
}

type Service struct {
	Candies
}

func New() *Service {
	return &Service{
		Candies: NewCandies(),
	}
}
