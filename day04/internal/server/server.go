package server

import (
	"day04/ex00/internal/transport"
	"day04/ex00/utils"
	"fmt"
	"log"
	"net/http"
)

const (
	certFile = "./cert/localhost/cert.pem"
	keyFile  = "./cert/localhost/key.pem"
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
	log.Println("Running HTTP server")
	return s.httpServer.ListenAndServe()
}

func (s *Server) RunTLS() error {
	log.Println("Running HTTPS server")
	s.httpServer.TLSConfig, _ = utils.GetTlsConfig(certFile, keyFile)
	return s.httpServer.ListenAndServeTLS("", "")
}

func (s *Server) Stop() error {
	return s.httpServer.Close()
}
