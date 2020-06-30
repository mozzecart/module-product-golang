package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	db "github.com/mozze/module-product-golang"
	"github.com/mozze/module-product-golang/models"
)

func fetch() ([]*models.Product, error) {
	var results []*models.Product

	ctx := context.Background()
	findOptions := options.Find()

	collection, errConnect := db.Product()
	if errConnect != nil {
		return nil, errConnect
	}

	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var elem models.Product
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, &elem)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	err = cur.Close(ctx)

	return results, err
}
