package service

import (
	movieapi "movieAPI"
	"movieAPI/pkg/repository"
)

type MovieService struct {
	repo repository.MovieList
}

func NewMovieService(repo repository.MovieList) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) Create(role int, list movieapi.MovieList) (int, error) {
	return s.repo.Create(role, list)
}

func (s *MovieService) ListMovies(order string) ([]movieapi.MovieList, error) {
	return s.repo.ListMovies(order)
}

func (s *MovieService) GetByName(movieName string) ([]movieapi.MovieList, error) {
	return s.repo.GetByName(movieName)
}

func (s *MovieService) Update(role, movId int, input movieapi.UpdateMovieListInput) error {
	if err := input.Validation(); err != nil {
		return err
	}
	return s.repo.Update(role, movId, input)
}

func (s *MovieService) Delete(role, movId int) error {
	return s.repo.Delete(role, movId)
}
