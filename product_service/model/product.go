package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string  `bson:"name,omitempty"`
	Description string  `bson:"description,omitempty"`
	Price       float32 `bson:"price,omitempty"`
	Stock       int32     `bson:"stock,omitempty"`
}
