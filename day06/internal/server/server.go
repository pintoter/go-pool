package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer      *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(handler http.Handler, host, port string) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", host, port),
			Handler: handler,
		},
		notify:          make(chan error),
		shutdownTimeout: 5 * time.Second,
	}
}

func (s *Server) Run() {
	go func() {
		s.notify <- s.httpServer.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
