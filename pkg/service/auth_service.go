package service

import (
	"crypto/sha1"
	"fmt"
	movieapi "movieAPI"
	"movieAPI/pkg/repository"
)

const (
	salt = "slfnoinrf90h390fnviofkl"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user movieapi.User) (int, error) {
	user.Password = passHash(user.Password)
	return s.repo.CreateUser(user)
}
func (s *AuthService) SignUser(username, password string) (int, error) {
	id, err := s.repo.SignUser(username, passHash(password))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func passHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
