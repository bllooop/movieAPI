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

func (s *ActorService) CreateActor(userRole string, list movieapi.ActorList) (int, error) {
	return s.repo.CreateActor(userRole, list)
}

func (s *ActorService) ListActors() ([]movieapi.ActorList, error) {
	return s.repo.ListActors()
}

func (s *ActorService) Delete(userRole string, actorId int) error {
	return s.repo.Delete(userRole, actorId)
}
func (s *ActorService) Update(userRole string, actorId int, input movieapi.UpdateActorListInput) error {
	if err := input.ValidationAct(); err != nil {
		return err
	}
	return s.repo.Update(userRole, actorId, input)
}
