package repositories

import (
	"context"
	"go-mongodb-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository struct {
	Collection *mongo.Collection
}

func NewPaymentRepository(client *mongo.Client) *PaymentRepository {
	return &PaymentRepository{
		Collection: client.Database(dbName).Collection("payments"),
	}
}

func (r *PaymentRepository) CreatePayment(payment models.Payment) (*models.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.Collection.InsertOne(ctx, payment)
	if err != nil {
		return nil, err
	}

	payment.ID = result.InsertedID.(primitive.ObjectID)
	return &payment, nil
}
