package handlers

import (
	"encoding/json"
	"fmt"
	movieapi "movieapi"
	"net/http"
	"strconv"
)

// @Summary Create movie list
// @Security ApiKeyAuth
// @Tags movieLists
// @Description create movie list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body movieapi.MovieList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {string} message
// @Failure 500 {string} message
// @Failure default {string} message
// @Router /api/movies/add [post]
func (h *Handler) createMovielist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		clientErr(w, http.StatusMethodNotAllowed, "only post method allowed")
		return
	}
	//retrievedValue := "1" // when testing uncomment
	retrievedValue := r.Context().Value(roleCtx).(string) // when testing comment
	var input movieapi.MovieList
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Title == "" || input.Date == "" || input.Description == "" || input.Rating == 0 || len(input.ActorName) == 0 {
		clientErr(w, http.StatusBadRequest, "invalid input body")
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

// @Summary Get all movies list
// @Security ApiKeyAuth
// @Tags movieLists
// @Description get all movies in list
// @ID get-list
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {string} message
// @Failure 500 {string} message
// @Failure default {string} message
// @Router /api/movies [get]
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

// @Summary Find movie in list
// @Security ApiKeyAuth
// @Tags movieLists
// @Description find movie in list either by fragment of a movie or an actor's name
// @ID find-list
// @Produce  json
// @Param       name    query     string  false  "name search by name"
// @Success 200 {integer} integer 1
// @Failure 400,404 {string} message
// @Failure 500 {string} message
// @Failure default {string} message
// @Router /api/movie [get]
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

// @Summary Update movie in list
// @Security ApiKeyAuth
// @Tags movieLists
// @Description update movie in list by id
// @ID update-list
// @Accept  json
// @Produce  json
// @Param input body movieapi.MovieList true "list info"
// @Param       id    query     int  false  "movie update by id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {string} message
// @Failure 500 {string} message
// @Failure default {string} message
// @Router /api/movies/update [post]
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

// @Summary Delete movie from list
// @Security ApiKeyAuth
// @Tags movieLists
// @Description delete movie from list by id
// @ID delete-list
// @Produce  json
// @Param       id    query     int  false  "movie delete by id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {string} message
// @Failure 500 {string} message
// @Failure default {string} message
// @Router /api/movies/delete [post]
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
