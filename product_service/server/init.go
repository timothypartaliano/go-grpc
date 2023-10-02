package server

import (
	pb "preview-w2/pb"
	"preview-w2/product_service/repository"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	Repository repository.ProductRepository
}

func NewProductsServer(r repository.ProductRepository) ProductServer {
	return ProductServer{Repository: r}
}