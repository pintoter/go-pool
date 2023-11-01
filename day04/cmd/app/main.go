package main

import (
	"day04/ex00/internal/app"
	"flag"
	"log"
)

func main() {
	var isSecure bool

	flag.BoolVar(&isSecure, "tls", false, "TLS server")
	flag.Parse()

	if err := app.Run(isSecure); err != nil {
		log.Fatal(err)
	}
}
