package service

import (
	movieapi "movieapi"
	"movieapi/pkg/repository"
)

type Authorization interface {
	CreateUser(user movieapi.User) (int, error)
	CreateToken(shopname, password string) (string, error)
	ParseToken(accesstok string) (string, error)
	//GetUserRole()
}
type MovieList interface {
	Create(userRole string, list movieapi.MovieList) (int, error)
	ListMovies(order string) ([]movieapi.MovieList, error)
	GetByName(movieName string) ([]movieapi.MovieList, error)
	Delete(userRole string, movId int) error
	Update(userRole string, movId int, input movieapi.UpdateMovieListInput) error
}
type ActorList interface {
	CreateActor(userRole string, list movieapi.ActorList) (int, error)
	ListActors() ([]movieapi.ActorList, error)
	Delete(userRole string, actorId int) error
	Update(userRole string, actorId int, input movieapi.UpdateActorListInput) error
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
