package main

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	db "github.com/mozze/module-product-golang"
)

func deleteByID(bodyID string) (*mongo.DeleteResult, error) {
	ctx := context.Background()

	collection, errConnect := db.Product()
	if errConnect != nil {
		log.Fatal(errConnect)
		return nil, errConnect
	}

	id, errParser := primitive.ObjectIDFromHex(bodyID)
	if errParser != nil {
		return nil, errors.New("invalid id")
	}

	result, err := collection.DeleteOne(ctx, bson.M{"id": id})

	return result, err
}
