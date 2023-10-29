package transport

import (
	"encoding/json"
	"log"
	"net/http"
)

type buyCandyResponse struct {
	Change  int64  `json:"change"`
	Message string `json:"thanks"`
}

type buyCandyErrorResponse struct {
	Error string `json:"error"`
}

const thank = "Thank you!"

func newResponseSuccess(w http.ResponseWriter, r *http.Request, statusCode int, change int64) {
	log.Printf("[%s] %s - SOLD", r.Method, r.URL.Path)

	resp, _ := json.Marshal(buyCandyResponse{
		Change:  change,
		Message: thank,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func newResponseError(w http.ResponseWriter, r *http.Request, statusCode int, err string) {
	log.Printf("[%s] %s - Error: %s", r.Method, r.URL.Path, err)

	resp, _ := json.Marshal(buyCandyErrorResponse{
		Error: err,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}
