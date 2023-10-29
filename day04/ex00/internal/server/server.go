package server

import (
	"day04/ex00/internal/transport"
	"fmt"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func New(host, port string) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", host, port),
			Handler: transport.NewHandler(),
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.httpServer.Close()
}
