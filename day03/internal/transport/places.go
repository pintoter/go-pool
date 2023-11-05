package transport

import (
	"Day03/internal/entity"
	"fmt"
	"net/http"
	"strconv"
)

const (
	maxPage = 1365
	minPage = 1
)

func (h *Handler) getPlacesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		query := r.URL.Query().Get("page")

		var pageNumber int = 1
		if query != "" {
			var err error
			pageNumber, err = strconv.Atoi(query)
			if err != nil || pageNumber < minPage || pageNumber > maxPage {
				newErrorResponse(w, r, http.StatusBadRequest, fmt.Sprintf(entity.InvalidQuery.Error(), query))
				return
			}
		}

		places, count, err := h.service.GetPlaces(10, (pageNumber-1)*10)
		if err != nil {
			newErrorResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		newPlacesWithTemplateResponse(w, places, count, pageNumber)
	} else {
		newErrorResponse(w, r, http.StatusBadRequest, "not implemented method")
	}
}

func (h *Handler) getPlacesJsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		query := r.URL.Query().Get("page")

		var pageNumber int = 1
		if query != "" {
			var err error
			pageNumber, err = strconv.Atoi(query)
			if err != nil || pageNumber < 1 || pageNumber > 1365 {
				newErrorResponse(w, r, http.StatusBadRequest, fmt.Sprintf(entity.InvalidQuery.Error(), query))
				return
			}
		}

		places, count, err := h.service.GetPlaces(10, (pageNumber-1)*10)
		if err != nil {
			newErrorResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		newPlacesResponse(w, places, count, pageNumber)
	} else {
		newErrorResponse(w, r, http.StatusBadRequest, "not implemented method")
	}
}

func (h *Handler) getClosestPlacesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryLat := r.URL.Query().Get("lat")
		queryLon := r.URL.Query().Get("lon")

		var lat, lon = 0.0, 0.0
		if queryLat != "" {
			var err error
			lat, err = strconv.ParseFloat(queryLat, 64)
			if err != nil || lat < 0 {
				newErrorResponse(w, r, http.StatusBadRequest, fmt.Sprintf(entity.InvalidQuery.Error(), queryLat))
				return
			}
		}

		if queryLon != "" {
			var err error
			lon, err = strconv.ParseFloat(queryLon, 64)
			if err != nil || lon < 0 {
				newErrorResponse(w, r, http.StatusBadRequest, fmt.Sprintf(entity.InvalidQuery.Error(), queryLon))
				return
			}
		}

		places, err := h.service.GetClosestPlaces(lat, lon)
		if err != nil {
			newErrorResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		newRecomendPlacesResponse(w, places)
	} else {
		newErrorResponse(w, r, http.StatusBadRequest, "not implemented method")
	}
}

func (h *Handler) getJwtTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		
	} else {
		newErrorResponse(w, r, http.StatusBadRequest, "not implemented method")
	}
}
