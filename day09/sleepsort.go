package main

import (
	"sync"
	"time"
)

func sleepSort(arr []int) chan int {
	ch := make(chan int)

	if len(arr) == 0 || arr == nil {
		close(ch)
		return ch
	}

	var wg sync.WaitGroup

	for _, value := range arr {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			time.Sleep(time.Duration(x) * time.Second)
			ch <- x
		}(value)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
