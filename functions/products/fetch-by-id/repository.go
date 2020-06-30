package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	db "github.com/mozze/module-product-golang"
	"github.com/mozze/module-product-golang/models"
)

func fetchByID(bodyID string) (*models.Product, error) {
	var result *models.Product

	id, errParser := primitive.ObjectIDFromHex(bodyID)
	if errParser != nil {
		return nil, errors.New("invalid id")
	}

	collection, err := db.Product()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	findResult := collection.FindOne(ctx, bson.D{{"id", id}})
	if err := findResult.Err(); err != nil {
		return nil, err
	}

	err = findResult.Decode(&result)

	return result, err
}
