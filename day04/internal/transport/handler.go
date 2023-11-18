package transport

import (
	"day04/internal/service"
	"log"
	"net/http"
)

type Handler struct {
	mux     *http.ServeMux
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	handler := &Handler{
		mux:     http.NewServeMux(),
		service: service,
	}

	handler.mux.HandleFunc("/buy_candy", func(w http.ResponseWriter, r *http.Request) {
		handler.buyCandyHandler(w, r)
	})

	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL)

	h.mux.ServeHTTP(w, r)
}
