package services

import (
	"go-mongodb-api/models"
	"go-mongodb-api/pkg/websocket"
	"go-mongodb-api/repositories"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	Repo    *repositories.UserRepository
	Manager *websocket.WebSocketManager
}

func NewUserService(repo *repositories.UserRepository, manager *websocket.WebSocketManager) *UserService {
	return &UserService{
		Repo:    repo,
		Manager: manager,
	}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.FetchAll()
}

func (s *UserService) GetUserByID(id primitive.ObjectID) (*models.User, error) {
	return s.Repo.FetchByID(id)
}

func (s *UserService) CreateUser(user models.User) (*models.User, error) {
	createdUser, err := s.Repo.Create(user)
	if err != nil {
		return nil, err
	}

	// goroutine
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Error during broadcast: %v", r)
			}
		}()
		message := "A new user was created. ID: " + createdUser.ID.Hex()
		s.Manager.BroadcastMessage(message)
	}()

	return createdUser, nil
}

func (s *UserService) UpdateUser(user models.User) (*models.User, error) {
	return s.Repo.Update(user)
}
