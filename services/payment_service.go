package services

import (
	"go-mongodb-api/models"
	"go-mongodb-api/repositories"
	// "log"
)

type PaymentService struct {
	Repo *repositories.PaymentRepository
}

func NewPaymentService(repo *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{
		Repo: repo,
	}
}

func (s *PaymentService) CreatePayment(payment models.Payment) (*models.Payment, error) {
	createdPayment, err := s.Repo.CreatePayment(payment)
	if err != nil {
		return nil, err
	}

	// message := "A new user was created. ID: " + createdPayment.ID.Hex()

	return createdPayment, nil
}
