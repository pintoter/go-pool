package main

import (
	"fmt"
	"sync"
)

func multiplex(chans ...<-chan interface{}) chan interface{} {
	res := make(chan interface{})
	var wg sync.WaitGroup

	for ch := range chans {
		wg.Add(1)
		go func(ch interface{}) {
			defer wg.Done()
			res <- ch
		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

const TotalGoroutines = 150_000

func main() {
	input1 := make(chan interface{})
	input2 := make(chan interface{})
	input3 := make(chan interface{})

	go func() {
		for i := 0; i < TotalGoroutines/3; i++ {
			input1 <- i
			input2 <- i * 10
			input3 <- i * 100
		}
		close(input1)
		close(input2)
		close(input3)
	}()

	fanIn := multiplex(input1, input2, input3)

	c := 0
	for data := range fanIn {
		c++
		fmt.Println(data)
	}
}
