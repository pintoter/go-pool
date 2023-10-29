package app

import (
	"day04/ex00/internal/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	host = "127.0.0.1"
	port = "3333"
)

func Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/buy_candy", loggingMiddleware(buyHandler))

	addr := fmt.Sprintf("%s:%s", host, port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		return err
	}
	return nil
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL)
		next(w, r)
	}
}

func buyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		buyCandy(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

type buyCandyInput struct {
	Money int64  `json:"money"`
	Type  string `json:"candyType"`
	Count int64  `json:"candyCount"`
}

type buyCandyResponse struct {
	Change  int64  `json:"change"`
	Message string `json:"thanks"`
}

type buyCandyErrorResponse struct {
	Error string `json:"error"`
}

var candyMenu = map[string]int64{
	"CE": 10,
	"AA": 15,
	"NT": 17,
	"DE": 21,
	"YR": 23,
}

func buyCandy(w http.ResponseWriter, r *http.Request) {
	var input buyCandyInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if _, ok := candyMenu[input.Type]; !ok || input.Count <= 0 || input.Money <= 0 {
		resp, _ := json.Marshal(buyCandyErrorResponse{
			Error: entity.ErrorInInputData,
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
	} else if input.Money < candyMenu[input.Type]*input.Count {
		value := fmt.Sprintf(entity.NotEnoughMoney, candyMenu[input.Type]*input.Count-input.Money)
		resp, _ := json.Marshal(buyCandyErrorResponse{
			Error: value,
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusPaymentRequired)
		w.Write(resp)
	} else {
		resp, _ := json.Marshal(buyCandyResponse{
			Change:  input.Money - candyMenu[input.Type]*input.Count,
			Message: "Thank you!",
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(resp)
	}
	w.Write([]byte("\n"))
}
