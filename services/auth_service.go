package services

import (
	"errors"
	"go-mongodb-api/repositories"
)

type AuthService struct {
	Repo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.Repo.FetchByEmail(username)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("invalid username or password")
	}

	return "sample-token", nil
}
