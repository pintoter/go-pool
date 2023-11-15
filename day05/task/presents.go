package task

import (
	"container/heap"
	cheap "day05/common/heap"
	"errors"
	"fmt"
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

func main() {
	t1 := []cheap.Present{
		{
			Value: 5,
			Size:  1,
		},
		{
			Value: 4,
			Size:  5,
		},
		{
			Value: 3,
			Size:  1,
		},
		{
			Value: 5,
			Size:  2,
		},
	}

	fmt.Println(getNCoolestPresents(t1, 2))
}
