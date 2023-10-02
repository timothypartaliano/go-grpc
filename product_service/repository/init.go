package repository

import "go.mongodb.org/mongo-driver/mongo"

type ProductRepository struct {
	DB *mongo.Database
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	return ProductRepository{DB: db}
}