package handlers

import (
	"encoding/json"
	"fmt"
	movieapi "movieAPI"
	"net/http"
)

func (h *Handler) createMovielist(w http.ResponseWriter, r *http.Request) {
	var input movieapi.MovieList
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		servErr(w, err)
		return
	}
	id, err := h.services.MovieList.Create(1, input)
	if err != nil {
		servErr(w, err)
		return
	}
	res, err := JSONStruct(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		servErr(w, err)
	}
	fmt.Fprintf(w, "%v", res)

}

func (h *Handler) getAllMoviesList(w http.ResponseWriter, r *http.Request) {
	order := r.URL.Query().Get("order")
	lists, err := h.services.ListMovies(order)
	if err != nil {
		servErr(w, err)
	}

	res, err := JSONStruct(lists)
	if err != nil {
		servErr(w, err)
	}
	fmt.Fprintf(w, "%v", res)
}

func (h *Handler) findMovieByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	list, err := h.services.MovieList.GetByName(name)
	if err != nil {
		servErr(w, err)
	}

	res, err := JSONStruct(list)
	if err != nil {
		servErr(w, err)
	}
	fmt.Fprintf(w, "%v", res)
}
