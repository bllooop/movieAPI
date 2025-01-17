package service

import (
	"errors"
	movieapi "movieapi"
	"movieapi/pkg/repository"
)

type MovieService struct {
	repo repository.MovieList
}

func NewMovieService(repo repository.MovieList) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) Create(userRole string, list movieapi.MovieList) (int, error) {
	return s.repo.Create(userRole, list)
}

func (s *MovieService) ListMovies(order string) ([]movieapi.MovieList, error) {
	return s.repo.ListMovies(order)
}

func (s *MovieService) GetByName(movieName string) ([]movieapi.MovieList, error) {
	return s.repo.GetByName(movieName)
}

func (s *MovieService) Update(userRole string, movId int, input movieapi.UpdateMovieListInput) error {
	if input.Title == nil && input.Description == nil && input.Rating == nil && input.Date == nil && input.ActorName == nil {
		return errors.New("update params have no values")
	}
	return s.repo.Update(userRole, movId, input)
}

func (s *MovieService) Delete(userRole string, movId int) error {
	return s.repo.Delete(userRole, movId)
}
