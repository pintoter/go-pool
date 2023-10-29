package transport

import (
	"day04/ex00/internal/service"
	"log"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler() *Handler {
	return &Handler{
		service: service.New(),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL)
	if r.Method == http.MethodPost {
		if r.URL.Path == "/buy_candy" {
			h.buyCandyHandler(w, r)
		} else {
			http.Error(w, "page isn't exist", http.StatusNotFound)
		}
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
