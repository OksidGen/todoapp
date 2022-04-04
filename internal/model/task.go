package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID     primitive.ObjectID `bson:"_id"`
	Text   string             `bson:"text"`
	Status bool               `bson:"status"`
}
