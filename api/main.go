package main

import (
	"log"
	"net/http"
	"preview-w2/api/controller"
	pbProduct "preview-w2/pb"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	defer conn.Close()
	
	// product client
	product := pbProduct.NewProductServiceClient(conn)

	// init controller
	productController := controller.NewProductController(product)

	// echo instance
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

	// create
	e.POST("/products",productController.CreateProduct)
	// view all
	e.GET("/products",productController.ViewProducts)
	// view by id
	e.GET("/products/:id",productController.ViewProductById)
	// update
	e.PUT("/products/:id",productController.UpdateProductById)
	// delete
	e.DELETE("/products/:id",productController.DeleteProductById)
	

    e.Logger.Fatal(e.Start(":8080"))
}