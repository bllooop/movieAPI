package service

import (
	"movieAPI/pkg/repository"
)

type MovieService struct {
	repo repository.Movies
}

func NewMovieService(repo repository.Movies) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) ListMovies() error {
	_, err := s.repo.ListMovies()
	if err != nil {
		return err
	}
	return nil
}
