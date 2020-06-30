package main

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	db "github.com/mozze/module-product-golang"
)

func updateByID(bodyID string, body *map[string]interface{}) (*mongo.UpdateResult, error) {
	ctx := context.Background()

	id, errParser := primitive.ObjectIDFromHex(bodyID)
	if errParser != nil {
		return nil, errors.New("invalid id")
	}

	collection, err := db.Product()
	if err != nil {
		return nil, err
	}

	(*body)["updated_at"] = time.Now()

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"id": id},
		bson.M{
			"$set": body,
		},
	)

	return result, err
}
