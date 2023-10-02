package server

import (
	"context"
	pb "preview-w2/pb"
	"preview-w2/product_service/model"
)

func (s ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error){
	pbProduct := pb.Product{}

	product := model.Product{
		Name: req.Product.Name,
		Description: req.Product.Description,
		Price: req.Product.Price,
		Stock: req.Product.Stock,
	}
	newProduct,err := s.Repository.CreateProduct(product)
	if err != nil {
		return &pbProduct, err
	}

	pbProduct.Name = newProduct.Name
	pbProduct.Description = newProduct.Description
	pbProduct.Price = newProduct.Price
	pbProduct.Stock = newProduct.Stock
	pbProduct.Id = newProduct.ID.Hex()


	return &pbProduct,nil
}

// func (pb.UnimplementedProductServiceServer).CreateProduct(context.Context, *pb.CreateProductRequest) (*pb.Product, error)
// func (pb.UnimplementedProductServiceServer).DeleteProduct(context.Context, *emptypb.Empty) (*pb.Product, error)
// func (pb.UnimplementedProductServiceServer).GetAllProduct(context.Context, *pb.GetAllProductRequest) (*pb.GetAllProductResponse, error)
// func (pb.UnimplementedProductServiceServer).GetProduct(context.Context, *pb.GetProductRequest) (*pb.Product, error)
// func (pb.UnimplementedProductServiceServer).UpdateProduct(context.Context, *pb.UpdateProductRequest) (*pb.Product, error)