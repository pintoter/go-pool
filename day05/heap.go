package main

import (
	"container/heap"
	"errors"
	"fmt"
)

type Present struct {
	Value int
	Size  int
}

type Presents []*Present

func (p Presents) Len() int { return len(p) }

func (p Presents) Less(i, j int) bool {
	if p[i].Value < p[j].Value {
		return true
	}

	if p[i].Value == p[j].Value && p[i].Size < p[j].Size {
		return true
	}

	return false
}

func (p Presents) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Presents) Push(x interface{}) {
	*p = append(*p, x.(*Present))
}

func (p *Presents) Pop() interface{} {
	old := *p
	n := old.Len()
	x := old[n-1]
	*p = old[:n-1]
	return x
}

func getNCoolestPresents(presents []Present, n int) ([]*Present, error) {
	if n > len(presents) || n <= 0 {
		return nil, errors.New("uncorrent N-size")
	}

	ph := make(Presents, n)
	for i, j := range presents {
		ph[i] = &Present{
			Value: j.Value,
			Size:  j.Size,
		}
	}

	heap.Init(&ph)

	coolestPresent := make([]*Present, n)
	for i := 0; i < n; n++ {
		coolestPresent[i] = heap.Pop(&ph).(*Present)
	}

	return coolestPresent, nil
}

func main() {
	presents := []Present{
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

	fmt.Println(getNCoolestPresents(presents, len(presents)))
}
