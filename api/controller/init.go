package controller

import "preview-w2/pb"

type ProductController struct {
	Service pb.ProductServiceClient
}

func NewProductController(pb pb.ProductServiceClient) ProductController {
	return ProductController{Service: pb}
}