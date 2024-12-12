package repositories

import (
	"context"
	"errors"
	"go-mongodb-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbName = "testdbgo"

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		Collection: client.Database(dbName).Collection("users"),
	}
}

func (r *UserRepository) FetchAll() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FetchByID(id primitive.ObjectID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FetchByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return &user, nil
}

func (r *UserRepository) Update(user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateFields := bson.M{}
	if user.Name != "" {
		updateFields["name"] = user.Name
	}
	if user.Email != "" {
		updateFields["email"] = user.Email
	}

	if user.Status != "" {
		updateFields["status"] = user.Status
	}

	if len(updateFields) == 0 {
		return nil, errors.New("filed are require")
	}

	update := bson.M{"$set": updateFields}
	filter := bson.M{"_id": user.ID}

	_, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
