package task

import (
	"container/heap"
	cheap "day05/common/heap"
	"errors"
)

func getNCoolestPresents(presents []cheap.Present, n int) ([]cheap.Present, error) {
	if n > len(presents) || n <= 0 {
		return nil, errors.New("invalid size")
	}

	ph := make(cheap.Presents, len(presents))
	for i, j := range presents {
		ph[i] = &cheap.Present{
			Value: j.Value,
			Size:  j.Size,
		}
	}

	heap.Init(&ph)

	coolestPresent := make([]cheap.Present, n)
	for i := 0; i < n; i++ {
		coolestPresent[i] = *(heap.Pop(&ph).(*cheap.Present))
	}

	return coolestPresent, nil
}
