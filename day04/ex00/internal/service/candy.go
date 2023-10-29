package service

import "day04/ex00/internal/entity"

var candyMenu = map[string]int64{
	"CE": 10,
	"AA": 15,
	"NT": 17,
	"DE": 21,
	"YR": 23,
}

type CandiesService struct{}

func NewCandies() *CandiesService {
	return &CandiesService{}
}

func (c *CandiesService) Buy(name string, count, money int64) (int64, error) {
	if _, ok := candyMenu[name]; !ok || count <= 0 || money <= 0 {
		return 0, entity.ErrorInInputData
	} else if money < candyMenu[name]*count {
		return candyMenu[name]*count - money, entity.NotEnoughMoney
	}

	return money - candyMenu[name]*count, nil
}
