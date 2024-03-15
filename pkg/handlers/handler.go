package handlers

import (
	"movieAPI/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/auth/sign-up", h.signUp)
	mux.HandleFunc("/auth/sign-in", h.signIn)
	mux.HandleFunc("/movies", h.getAllMoviesList)
	mux.HandleFunc("/movies/add", h.createMovielist)
	mux.HandleFunc("/movies/update", h.updateMovieList)
	mux.HandleFunc("/movies/delete", h.deleteMovieList)
	mux.HandleFunc("/actors", h.getAllActorList)
	mux.HandleFunc("/actors/add", h.createActorlist)
	mux.HandleFunc("/movie", h.findMovieByName)
	mux.HandleFunc("/actors/update", h.updateActorList)
	mux.HandleFunc("/actors/delete", h.deleteActorList)
	return mux
}
