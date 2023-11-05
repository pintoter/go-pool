package transport

import (
	"Day03/internal/service"
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

	handler.mux.HandleFunc("/places", func(w http.ResponseWriter, r *http.Request) {
		handler.getPlacesHandler(w, r)
	})

	handler.mux.HandleFunc("/api/places", func(w http.ResponseWriter, r *http.Request) {
		handler.getPlacesJsonHandler(w, r)
	})

	handler.mux.HandleFunc("/api/recommend", func(w http.ResponseWriter, r *http.Request) {
		handler.getClosestPlacesHandler(w, r)
	})

	handler.mux.HandleFunc("/api/get_token", func(w http.ResponseWriter, r *http.Request) {
		handler.getJwtTokenHandler(w, r)
	})

	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL)

	h.mux.ServeHTTP(w, r)
}
