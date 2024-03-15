package repository

import (
	"database/sql"
	movieapi "movieAPI"
)

type Authorization interface {
	CreateUser(user movieapi.User) (int, error)
	SignUser(username, password string) (int, error)
}
type MovieList interface {
	Create(role int, list movieapi.MovieList) (int, error)
	ListMovies(order string) ([]movieapi.MovieList, error)
	GetByName(movieName string) ([]movieapi.MovieList, error)
	Update(role, movId int, input movieapi.UpdateMovieListInput) error
	Delete(role, movId int) error
}

type ActorList interface {
	CreateActor(role int, list movieapi.ActorList) (int, error)
	ListActors() ([]movieapi.ActorList, error)
	Update(role, actorId int, input movieapi.UpdateActorListInput) error
	Delete(role, actorId int) error
}
type Repository struct {
	Authorization
	MovieList
	ActorList
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		MovieList:     NewMoviePostgres(db),
		ActorList:     NewActorPostgres(db),
	}
}
