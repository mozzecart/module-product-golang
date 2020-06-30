package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"id"`
	Category    []string           `json:"category"`
	Description string             `json:"description"`
	Gender      string             `json:"gender"`
	MadeIn      string             `json:"madein"`
	Name        string             `json:"name"`
	Price       float32            `json:"price"`
	Slug        string             `json:"slug"`
	Variations  []ProductVariation `json:"variations"`
	Visible     bool               `json:"visible"`

	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type ProductVariation struct {
	Color `json:"color"`
	Image string   `json:"image"`
	Size  []string `json:"size"`
	Stock int      `json:"stock"`
}
