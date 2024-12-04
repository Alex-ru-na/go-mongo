package services

import (
	"go-mongodb-api/models"
	"go-mongodb-api/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.FetchAll()
}

func (s *UserService) GetUserByID(id primitive.ObjectID) (*models.User, error) {
	return s.Repo.FetchByID(id)
}

func (s *UserService) CreateUser(user models.User) (*models.User, error) {
	return s.Repo.Create(user)
}

func (s *UserService) UpdateUser(user models.User) (*models.User, error) {
	return s.Repo.Update(user)
}
