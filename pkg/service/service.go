package service

import (
	movieapi "movieAPI"
	"movieAPI/pkg/repository"
)

type Authorization interface {
	CreateUser(user movieapi.User) (int, error)
	SignUser(username, password string) (int, error)
}
type Movies interface {
	ListMovies() error
}
type Service struct {
	Authorization
	Movies
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Movies:        NewMovieService(repos.Movies),
	}
}
