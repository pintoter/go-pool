package tools

import (
	"Day02/internal/entity"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type WcFlags struct {
	l, m, w bool
}

var wc WcFlags

func init() {
	flag.BoolVar(&wc.l, "l", false, "counting lines")
	flag.BoolVar(&wc.m, "m", false, "counting characters")
	flag.BoolVar(&wc.w, "w", false, "counting words")
}

func RunWc() {
	flag.Parse()

	err := checkFlags()
	if err != nil {
		log.Fatal(err)
	}

	files := flag.Args()

	if len(files) == 0 {
		log.Fatal(entity.ArgumentIsEmpty)
	}

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go counter(file, &wg)
	}

	wg.Wait()
}

func checkFlags() error {
	if (wc.l && wc.m) || (wc.l && wc.w) || (wc.m && wc.w) {
		return entity.TooMuchFlags
	}

	if !wc.l && !wc.m && !wc.w {
		wc.w = true
	}

	return nil
}

func counter(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fi, err := os.Stat(name)
	if err != nil {
		log.Println(err)
		return
	}

	if fi.Mode().IsRegular() {
		file, err := os.Open(name)
		if err != nil {
			log.Println(err)
			return
		}

		defer file.Close()

		var lines, words, bytes int

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			lines++

			line := scanner.Text()
			bytes += len(line)

			splitLines := strings.Fields(line)
			words += len(splitLines)
		}
		bytes += lines

		if wc.l {
			fmt.Printf("%d\t%s\n", lines, name)
		} else if wc.m {
			fmt.Printf("%d\t%s\n", bytes, name)
		} else if wc.w {
			fmt.Printf("%d\t%s\n", words, name)
		}

	}
}
