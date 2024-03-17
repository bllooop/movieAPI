package handlers

import (
	"encoding/json"
	"fmt"
	movieapi "movieapi"
	"net/http"
	"strconv"
)

func (h *Handler) createMovielist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		clientErr(w, http.StatusMethodNotAllowed, "only post method allowed")
		return
	}
	retrievedValue := r.Context().Value(roleCtx).(string)
	var input movieapi.MovieList
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		clientErr(w, http.StatusBadRequest, err.Error())
		return
	}
	if input.Rating > 10 || input.Rating < 0 {
		clientErr(w, http.StatusBadRequest, "rating allowed between 0 and 10")
	}
	id, err := h.services.MovieList.Create(retrievedValue, input)
	if err != nil {
		servErr(w, err, err.Error())
		return
	}
	res, err := JSONStruct(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		servErr(w, err, err.Error())
	}
	fmt.Fprintf(w, "%v", res)
}

func (h *Handler) getAllMoviesList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/movies" {
		notFound(w)
		return
	}
	order := r.URL.Query().Get("order")
	lists, err := h.services.ListMovies(order)
	if err != nil {
		servErr(w, err, err.Error())
	}

	res, err := JSONStruct(lists)
	if err != nil {
		servErr(w, err, err.Error())
	}
	fmt.Fprintf(w, "%v", res)
}

func (h *Handler) findMovieByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	list, err := h.services.MovieList.GetByName(name)
	if err != nil {
		clientErr(w, http.StatusBadRequest, err.Error())
	}
	res, err := JSONStruct(list)
	if err != nil {
		servErr(w, err, err.Error())
	}
	fmt.Fprintf(w, "%v", res)
}
func (h *Handler) updateMovieList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		clientErr(w, http.StatusMethodNotAllowed, "only post method allowed")
		return
	}
	retrievedValue := r.Context().Value(roleCtx).(string)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		clientErr(w, http.StatusBadRequest, "invalid id parameter")
		return
	}
	var input movieapi.UpdateMovieListInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		clientErr(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.MovieList.Update(retrievedValue, id, input); err != nil {
		servErr(w, err, err.Error())
	}
	res, err := JSONStruct(statusResponse{
		Status: "ok",
	})
	if err != nil {
		servErr(w, err, err.Error())
	}
	fmt.Fprintf(w, "%v", res)
}

func (h *Handler) deleteMovieList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		clientErr(w, http.StatusMethodNotAllowed, "only delete method allowed")
		return
	}
	retrievedValue := r.Context().Value(roleCtx).(string)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		clientErr(w, http.StatusBadRequest, "invalid id parameter")
		return
	}
	err = h.services.MovieList.Delete(retrievedValue, id)
	if err != nil {
		servErr(w, err, err.Error())
		return
	}
	res, err := JSONStruct(statusResponse{
		Status: "ok",
	})
	if err != nil {
		servErr(w, err, err.Error())
	}
	fmt.Fprintf(w, "%v", res)
}
