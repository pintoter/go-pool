package transport

import (
	"Day03/internal/entity"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type getPlacesWithTemplateResponse struct {
	PageNum  int
	PrevPage int
	NextPage int
	Total    int
	Places   []entity.Place
}

type getPlacesJsonResponse struct {
	Name     string         `json:"name"`
	Total    int            `json:"total"`
	Places   []entity.Place `json:"places"`
	PrevPage int            `json:"prev_page,omitempty"`
	NextPage int            `json:"next_page,omitempty"`
	LastPage int            `json:"last_page"`
}

type getRecomendPlacesResponse struct {
	Name   string         `json:"name"`
	Places []entity.Place `json:"places"`
}

type getTokenResponse struct {
	Token string `json:"token"`
}

type getPlacesErrorResponse struct {
	Error string `json:"error"`
}

func newPlacesWithTemplateResponse(w http.ResponseWriter, places []entity.Place, count, pageNumber int) {
	var data getPlacesWithTemplateResponse
	data.PageNum = pageNumber
	data.PrevPage = pageNumber - 1
	data.NextPage = pageNumber + 1
	data.Total = count
	data.Places = places

	template, err := template.ParseFiles("./web/templates/templates.html")
	if err != nil {
		log.Println("error with parsing template", err)
	}

	template.Execute(w, data)
}

func newErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	log.Printf("[%s] %s - Response: Error: %s", r.Method, r.URL.Path, message)

	resp, _ := json.MarshalIndent(getPlacesErrorResponse{
		Error: message,
	}, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func newPlacesResponse(w http.ResponseWriter, places []entity.Place, count, pageNumber int) {
	var data getPlacesJsonResponse
	data.Name = "Places"
	data.Total = count
	data.Places = places
	data.LastPage = maxPage

	if pageNumber != minPage {
		data.PrevPage = pageNumber - 1
	}

	if pageNumber != maxPage {
		data.NextPage = pageNumber + 1
	}

	resp, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println("problem with marshalling places", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func newRecomendPlacesResponse(w http.ResponseWriter, places []entity.Place) {
	var data getRecomendPlacesResponse
	data.Name = "Recommendation"
	data.Places = places

	resp, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println("problem with marshalling places", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func newTokenResponse(w http.ResponseWriter, token string) {
	var data getTokenResponse
	data.Token = token

	resp, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println("problem with marshalling places", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
