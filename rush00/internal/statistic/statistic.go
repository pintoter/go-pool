package statistic

import "github.com/montanaflynn/stats"

type Statistic struct {
	Count      int64
	Mean       float64
	SD         float64
	Frequencys []float64
}

func NewStat() *Statistic {
	return &Statistic{
		Mean:       0,
		SD:         0,
		Frequencys: make([]float64, 0),
	}
}

func (s *Statistic) Update(value float64) {
	s.Count++
	s.Frequencys = append(s.Frequencys, value)
	s.Mean, _ = stats.Mean(s.Frequencys)
	s.SD, _ = stats.StandardDeviation(s.Frequencys)
}
