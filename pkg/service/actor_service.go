package service

import (
	movieapi "movieapi"
	"movieapi/pkg/repository"
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

func (s *ActorService) Delete(role, actorId int) error {
	return s.repo.Delete(role, actorId)
}
func (s *ActorService) Update(role, actorId int, input movieapi.UpdateActorListInput) error {
	if err := input.ValidationAct(); err != nil {
		return err
	}
	return s.repo.Update(role, actorId, input)
}
