package service

import (
	movieapi "movieAPI"
	"movieAPI/pkg/repository"
)

type Authorization interface {
	CreateUser(user movieapi.User) (int, error)
	SignUser(username, password string) (int, error)
}
type MovieList interface {
	Create(role int, list movieapi.MovieList) (int, error)
	ListMovies(order string) ([]movieapi.MovieList, error)
	GetByName(movieName string) ([]movieapi.MovieList, error)
	Delete(role, movId int) error
	Update(role, movId int, input movieapi.UpdateMovieListInput) error
}
type ActorList interface {
	CreateActor(role int, list movieapi.ActorList) (int, error)
	ListActors() ([]movieapi.ActorList, error)
	Delete(role, actorId int) error
	Update(role, actorId int, input movieapi.UpdateActorListInput) error
}
type Service struct {
	Authorization
	MovieList
	ActorList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		MovieList:     NewMovieService(repos.MovieList),
		ActorList:     NewActorService(repos.ActorList),
	}
}
