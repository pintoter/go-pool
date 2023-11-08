package main

import (
	"day04/internal/app"
	"day04/internal/config"
	"flag"
)

var isSecure bool

func init() {
	flag.BoolVar(&isSecure, "tls", false, "TLS server")
	flag.Parse()
}

func main() {
	cfg := config.Get()

	app.Run(cfg, isSecure)
}
