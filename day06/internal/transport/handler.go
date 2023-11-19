package transport

import (
	"context"
	"day06/internal/config"
	"day06/internal/service"
	"html/template"
	"log"
	"net/http"
	"strings"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Limit(100), 1)

type Handler struct {
	credentials *config.AdminCredentials
	mux         *http.ServeMux
	service     *service.Service
}

func NewHandler(config *config.AdminCredentials, service *service.Service) *Handler {
	handler := &Handler{
		mux:         http.NewServeMux(),
		service:     service,
		credentials: config,
	}

	img := http.FileServer(http.Dir("./web/image"))
	handler.mux.Handle("/img", http.StripPrefix("/img", img))

	md := http.FileServer(http.Dir("./web/md"))
	handler.mux.Handle("/md", http.StripPrefix("/md", md))

	handler.mux.HandleFunc("/admin", handler.AdminHandler)

	handler.mux.HandleFunc("/", handler.ArticlesHandler)
	handler.mux.HandleFunc("/article1.html", handler.ArticleHandler)

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

type Data struct {
	login, password string
}

func (h *Handler) AdminHandler(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	if login == h.credentials.Login && password == h.credentials.Password {
		http.ServeFile(w, r, "./ui/html/admin.html")
		articleName := r.FormValue("articleName")
		articleLink := r.FormValue("articleLink")
		if articleName != "" && articleLink != "" && strings.Contains(articleLink, ".md") {
			h.service.CreateArticle(context.Background(), articleName, articleLink)
		}
		log.Println("logged in as admin")
	}

	template.Must(template.ParseFiles("./ui/html/authentication.html")).Execute(w, Data{login, password})
}

func (h *Handler) ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	// login := r.FormValue("login")
	// password := r.FormValue("password")

	// if login == h.credentials.Login && password == h.credentials.Password {
	// 	http.ServeFile(w, r, "./ui/html/admin.html")
	// 	articleName := r.FormValue("articleName")
	// 	articleLink := r.FormValue("articleLink")
	// 	if articleName != "" && articleLink != "" && strings.Contains(articleLink, ".md") {
	// 		h.service.CreateArticle(context.Background(), articleName, articleLink)
	// 	}
	// 	log.Println("logged in as admin")
	// }

	// template.Must(template.ParseFiles("./ui/html/authentication.html")).Execute(w, Data{login, password})
}

func (h *Handler) ArticleHandler(w http.ResponseWriter, r *http.Request) {
	// login := r.FormValue("login")
	// password := r.FormValue("password")

	// if login == h.credentials.Login && password == h.credentials.Password {
	// 	http.ServeFile(w, r, "./ui/html/admin.html")
	// 	articleName := r.FormValue("articleName")
	// 	articleLink := r.FormValue("articleLink")
	// 	if articleName != "" && articleLink != "" && strings.Contains(articleLink, ".md") {
	// 		h.service.CreateArticle(context.Background(), articleName, articleLink)
	// 	}
	// 	log.Println("logged in as admin")
	// }

	// template.Must(template.ParseFiles("./ui/html/authentication.html")).Execute(w, Data{login, password})
}
