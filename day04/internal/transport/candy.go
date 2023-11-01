package transport

import (
	"day04/ex00/internal/entity"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const buyCandyPath = "/buy_candy"

type buyCandyInput struct {
	Money int64  `json:"money"`
	Type  string `json:"candyType"`
	Count int64  `json:"candyCount"`
}

func (h *Handler) buyCandyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if r.URL.Path == buyCandyPath {
			var input buyCandyInput

			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			money, err := h.service.Buy(input.Type, input.Count, input.Money)
			if err != nil {
				if errors.Is(err, entity.ErrorInInputData) {
					newResponseError(w, r, http.StatusBadRequest, err.Error())
				} else if errors.Is(err, entity.NotEnoughMoney) {
					newResponseError(w, r, http.StatusPaymentRequired, fmt.Sprintf(err.Error(), money))
				}
				return
			}

			newResponseSuccess(w, r, http.StatusCreated, money)
		} else {
			http.Error(w, "page is not exist", http.StatusNotFound)
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
