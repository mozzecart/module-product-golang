package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	db "github.com/mozze/module-product-golang"
	"github.com/mozze/module-product-golang/models"
)

func fetch() ([]*models.Category, error) {
	var results []*models.Category

	findOptions := options.Find()

	collection, errConnect := db.Category()
	if errConnect != nil {
		log.Fatal(errConnect)
		return nil, errConnect
	}

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem models.Category
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		results = append(results, &elem)
	}

	if err = cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = cur.Close(context.TODO()); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return results, nil
}
