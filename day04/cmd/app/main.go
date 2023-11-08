package main

import (
	"day04/internal/app"
	"flag"
)

var isSecure bool

func init() {
	flag.BoolVar(&isSecure, "tls", false, "TLS server")
	flag.Parse()
}

func main() {
	app.Run(isSecure)
}
