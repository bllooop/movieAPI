package handlers

import (
	"movieapi/pkg/service"
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
	mux.HandleFunc("/api/auth/sign-up", h.signUp)
	mux.HandleFunc("/api/auth/sign-in", h.signIn)
	mux.HandleFunc("/api/movies", h.AuthMiddleware(h.getAllMoviesList))
	mux.HandleFunc("/api/movies/add", h.AuthMiddleware(h.createMovielist))
	mux.HandleFunc("/api/movies/update", h.AuthMiddleware(h.updateMovieList))
	mux.HandleFunc("/api/movies/delete", h.AuthMiddleware(h.deleteMovieList))
	mux.HandleFunc("/api/actors", h.AuthMiddleware(h.getAllActorList))
	mux.HandleFunc("/api/actors/add", h.AuthMiddleware(h.createActorlist))
	mux.HandleFunc("/api/movie", h.AuthMiddleware(h.findMovieByName))
	mux.HandleFunc("/api/actors/update", h.AuthMiddleware(h.updateActorList))
	mux.HandleFunc("/api/actors/delete", h.AuthMiddleware(h.deleteActorList))
	//wrappedMux := h.authCheck(mux)
	return mux
}
