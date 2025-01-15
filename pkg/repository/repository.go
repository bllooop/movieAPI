package repository

import (
	"database/sql"
	movieapi "movieapi"
)

type Authorization interface {
	CreateUser(user movieapi.User) (int, error)
	SignUser(username, password string) (movieapi.User, error)
	// GetUserRole()
}
type MovieList interface {
	Create(userRole string, list movieapi.MovieList) (int, error)
	ListMovies(order string) ([]movieapi.MovieList, error)
	GetByName(movieName string) ([]movieapi.MovieList, error)
	Update(userRole string, movId int, input movieapi.UpdateMovieListInput) error
	Delete(userRole string, movId int) error
}

type ActorList interface {
	CreateActor(userRole string, list movieapi.ActorList) (int, error)
	ListActors() ([]movieapi.ActorList, error)
	Update(userRole string, actorId int, input movieapi.UpdateActorListInput) error
	Delete(userRole string, actorId int) error
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
