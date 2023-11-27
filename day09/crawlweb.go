package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const maxGoroutines = 8

func crawlWeb(ctx context.Context, urls <-chan string) chan *string {
	result := make(chan *string)
	ch := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup

	go func() {
		for url := range urls {
			select {
			case <-ctx.Done():
				fmt.Println("context done")
				break
			case ch <- struct{}{}:
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					defer func() {
						<-ch
					}()
					resp, err := http.Get(url)
					if err != nil {
						result <- new(string)
						return
					}
					defer resp.Body.Close()

					body, err := io.ReadAll(resp.Body)
					if err != nil {
						result <- new(string)
						return
					}
					output := string(body)
					result <- &output

				}(url)
			}
		}
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	var urls []string = []string{
		"bitcoin",
		"etherium",
		"tether",
		"bnb",
		"xrp",
		"usdc",
		"solana",
		"cardano",
		"dogecoin",
		"tron",
		"avalanche",
		"polkadot",
		"litecoin",
		"shibainu",
	}

	urlsCh := make(chan string)

	go func() {
		for _, url := range urls {
			urlsCh <- fmt.Sprintf("https://api.coincap.io/v2/assets/%s", url)
		}
		close(urlsCh)
	}()

	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGKILL, syscall.SIGTERM)

	go func() {
		s := <-sig
		fmt.Printf("got signal %v, cancel", s)
		cancel()
	}()

	res := crawlWeb(ctx, urlsCh)

	for body := range res {
		fmt.Println("GOTCHA")
		fmt.Println(*body)
	}
}
