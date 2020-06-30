package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(
		os.Getenv("MONGODB_URI"),
	)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Couldn't connect to the database: %v", err.Error()))
	} else {
		log.Println("Connected!")
	}

	return client.Database("mozze"), nil
}

func Product() (*mongo.Collection, error) {
	connection, err := Connect()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return connection.Collection("products"), nil
}

func Category() (*mongo.Collection, error) {
	connection, err := Connect()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return connection.Collection("categories"), nil
}
