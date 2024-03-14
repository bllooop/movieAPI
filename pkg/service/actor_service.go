package service

import (
	movieapi "movieAPI"
	"movieAPI/pkg/repository"
)

type ActorService struct {
	repo repository.ActorList
}

func NewActorService(repo repository.ActorList) *ActorService {
	return &ActorService{repo: repo}
}

func (s *ActorService) CreateActor(role int, list movieapi.ActorList) (int, error) {
	return s.repo.CreateActor(role, list)
}

func (s *ActorService) ListActors() ([]movieapi.ActorList, error) {
	return s.repo.ListActors()
}
