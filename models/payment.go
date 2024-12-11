package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderID primitive.ObjectID `bson:"orderId,omitempty" json:"orderId"`
	Detail  string             `bson:"detail" json:"detail"`
	Amount  float32            `bson:"amount" json:"amount"`
}
