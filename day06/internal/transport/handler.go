package transport

import (
	"day06/internal/service"
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Limit(100), 1)

type Handler struct {
	mux     *http.ServeMux
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	handler := &Handler{
		mux:     http.NewServeMux(),
		service: service,
	}

	img := http.FileServer(http.Dir("./web/image"))
	handler.mux.Handle("/img", http.StripPrefix("/img", img))

	md := http.FileServer(http.Dir("./web/md"))
	handler.mux.Handle("/md", http.StripPrefix("/md", md))

	handler.mux.HandleFunc("/admin", handler.AdminHandler)

	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL)

	if limiter.Allow() {
		h.mux.ServeHTTP(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
	}
}

func (h *Handler) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/article1.html" {
		http.ServeFile(w, r, "./ui/html/article1.html")
	}
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./ui/html/articles.html")
	}
}

func (h *Handler) AdminHandler(w http.ResponseWriter, r *http.Request) {

}
