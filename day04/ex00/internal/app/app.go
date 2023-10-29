package app

import (
	"day04/ex00/internal/server"
)

const (
	host = "127.0.0.1"
	port = "3333"
)

func Run() error {
	server := server.New(host, port)
	if err := server.Run(); err != nil {
		return err
	}
	return nil
}
