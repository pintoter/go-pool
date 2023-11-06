package transport

import (
	"Day03/internal/entity"
	"Day03/internal/service"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const (
	SECRET = "secret_key"
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
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			newErrorResponse(w, r, http.StatusUnauthorized, entity.UnauthorizedToken.Error())
		} else {
			jwtToken := authHeader[1]
			token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(SECRET), nil
			})

			if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				handler.getClosestPlacesHandler(w, r)
			} else {
				newErrorResponse(w, r, http.StatusUnauthorized, entity.UnauthorizedToken.Error())
			}
		}
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
