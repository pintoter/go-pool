package task

import (
	"container/heap"
	"day05/common"
	"errors"
)

func getNCoolestPresents(presents []common.Present, n int) ([]common.Present, error) {
	if n > len(presents) || n <= 0 {
		return nil, errors.New("invalid size")
	}

	ph := make(common.Presents, n)
	for i, j := range presents {
		ph[i] = &common.Present{
			Value: j.Value,
			Size:  j.Size,
		}
	}

	heap.Init(&ph)

	coolestPresent := make([]common.Present, n)
	for i := 0; ph.Len() > 0; i++ {
		coolestPresent[i] = *(heap.Pop(&ph).(*common.Present))
	}

	return coolestPresent, nil
}
