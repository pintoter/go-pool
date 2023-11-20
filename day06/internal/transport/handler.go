package transport

import (
	"day06/internal/config"
	"day06/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/time/rate"
)

type Handler struct {
	sessionStore *sessions.CookieStore
	credentials  *config.AdminCredentials
	router       *mux.Router
	service      *service.Service
}

func NewHandler(config *config.AdminCredentials, service *service.Service) *Handler {
	handler := &Handler{
		sessionStore: sessions.NewCookieStore([]byte("secret-key")),
		router:       mux.NewRouter(),
		service:      service,
		credentials:  config,
	}

	handler.router.PathPrefix("/web/image/").Handler(http.StripPrefix("/web/image/", http.FileServer(http.Dir("image"))))
	handler.router.HandleFunc("/", handler.articlesHandler).Methods("GET")
	handler.router.HandleFunc("/article/{id}", handler.articleHandler).Methods("GET")
	handler.router.HandleFunc("/login", handler.loginHandler)
	handler.router.HandleFunc("/admin", handler.adminHandler).Methods("GET")
	handler.router.HandleFunc("/admin", handler.createArticleHandler).Methods("POST")

	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL)

	limiter := rate.NewLimiter(rate.Limit(100), 1)

	if limiter.Allow() {
		h.router.ServeHTTP(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
	}
}
