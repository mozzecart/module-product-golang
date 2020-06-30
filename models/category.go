package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"id"`
	Name string             `json:"name"`
}
