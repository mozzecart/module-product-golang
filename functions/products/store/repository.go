package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	db "github.com/mozze/module-product-golang"
	"github.com/mozze/module-product-golang/models"
)

func store(body *models.Product) (*models.Product, error) {
	body.ID = primitive.NewObjectID()
	body.CreatedAt = time.Now()
	body.UpdatedAt = time.Now()

	collection, errConnect := db.Product()
	if errConnect != nil {
		return nil, errConnect
	}

	ctx := context.Background()
	_, err := collection.InsertOne(ctx, body)

	return body, err
}
