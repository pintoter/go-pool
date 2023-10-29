package app

import (
	"day04/ex00/internal/server"
	"log"
)

const (
	host = "127.0.0.1"
	port = "3333"
)

func Run(isSecure bool) error {
	server := server.New(host, port)
	if !isSecure {
		if err := server.Run(); err != nil {
			return err
		}
	} else if isSecure {
		if err := server.RunTLS(); err != nil {
			return err
		}
		log.Println("Starting HTTPS server")
	}
	return nil
}
