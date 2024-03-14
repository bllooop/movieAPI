package repository

import (
	"database/sql"
	movieapi "movieAPI"
)

type Authorization interface {
	CreateUser(user movieapi.User) (int, error)
	SignUser(username, password string) (int, error)
}
type Movies interface {
	ListMovies() (movieapi.User, error)
}

type Repository struct {
	Authorization
	Movies
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Movies:        NewMoviePostgres(db),
	}
}
