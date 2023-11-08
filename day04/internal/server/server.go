package server

import (
	"context"
	"day04/utils"
	"fmt"
	"log"
	"net/http"
)

const (
	certFile = "./config/cert/localhost/cert.pem"
	keyFile  = "./config/cert/localhost/key.pem"
)

type Server struct {
	httpServer *http.Server
}

func New(handler http.Handler, host, port string) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", host, port),
			Handler: handler,
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

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
