package main

import (
	"log"
	"net"
	"preview-w2/config"
	"preview-w2/pb"
	"preview-w2/product_service/repository"
	"preview-w2/product_service/server"

	"google.golang.org/grpc"
)

func main() {
	// mongodb
	db := config.InitMongoDB()

	// server
	srv := grpc.NewServer()

	// repository product
	productRepository := repository.NewProductRepository(db)

	// server product
	productServer := server.NewProductsServer(productRepository)
	PORT := ":50051"

	// register product service server
	pb.RegisterProductServiceServer(srv,productServer)
	log.Println("Starting RPC server at", PORT)

	// init listener
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", PORT, err)
	}

	log.Fatal(srv.Serve(listener))
}