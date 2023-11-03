package app

import (
	"Day03/internal/entity"
	"Day03/internal/repository"
	"Day03/pkg/database"
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	ESUsername    = "elastic"
	ESPW          = "Wqy0O9nLa8HCHZ*7MHLA"
	ESFingerPrint = "7140373329f32249e4775f12a8d4e35f5f2b2c390f4677aaf4f4b7c48c469ee9"
	address       = "http://localhost:9200"
	serveradr     = "127.0.0.1:8888"
)

type responsePlaces struct {
	PageNum  int
	PrevPage int
	NextPage int
	Total    int64
	Places   []entity.Place
}

func Run() {
	ctx := context.Background()

	es, err := database.New(address)
	if err != nil {
		log.Fatal(err)
	}

	// Check ES for existing index
	database.LoadData(ctx, es)
	err = database.UpdateMaxResultSettings(es)
	if err != nil {
		log.Println("error with update settings to 20k", err)
	}

	repo := repository.NewPlacesRepository(es)

	mux := http.NewServeMux()

	mux.HandleFunc("/places", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("page")

		var pageNumber int = 1
		if query != "" {
			pageNumber, err := strconv.Atoi(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if pageNumber < 1 || pageNumber > 1365 {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		places, count, err := repo.GetPlaces(10, (pageNumber-1)*10)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var resp responsePlaces
		resp.PageNum = pageNumber
		resp.PrevPage = pageNumber - 1
		resp.NextPage = pageNumber + 1
		resp.Total = int64(count)
		resp.Places = places

		template, err := template.ParseFiles("./web/templates/templates.html")
		if err != nil {
			log.Println("error with parsing template", err)
		}

		template.Execute(w, resp)
	})

	err = http.ListenAndServe(serveradr, mux)
	if err != nil {
		log.Println("error when runnig server", err)
	}

	// html/template
}

// TODO: 1) total 13650
// 2) порядок выдачи результатов
