package repository

import (
	"context"
	"preview-w2/product_service/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	// get coll
	result,err := r.DB.Collection("products").InsertOne(context.Background(),product)
	if err != nil {
		return model.Product{},err
	}

	// insert id to model
	product.ID = result.InsertedID.(primitive.ObjectID)
	return product, nil
}
