package controller

import (
	"preview-w2/pb"

	"github.com/labstack/echo/v4"
)

// create
func (c *ProductController) CreateProduct(ctx echo.Context) error{
	// bind
	reqBody := ReqBodyCreateProduct{}
	err := ctx.Bind(&reqBody)
	if err != nil {
		return echo.NewHTTPError(400,echo.Map{
			"message":"error to bind",
			"detail":err.Error(),
		})
	}

	// create
	req := pb.CreateProductRequest{
		Product: &pb.Product{
			Name: reqBody.Name,
			Description: reqBody.Description,
			Price: reqBody.Price,
			Stock: int32(reqBody.Stock),
		},
	}

	newProduct,err := c.Service.CreateProduct(ctx.Request().Context(),&req)
	if err != nil {
		return echo.NewHTTPError(400,echo.Map{
			"message":"error to create",
			"detail":err.Error(),
		})
	}

	return ctx.JSON(201,echo.Map{
		"message":"success create",
		"detail":newProduct,
	})
}

// view
func (c *ProductController) ViewProducts(ctx echo.Context) error{
	return ctx.JSON(200,echo.Map{
		"message":"success view products",
		"detail":"-",
	})
}

// view by id
func (c *ProductController) ViewProductById(ctx echo.Context) error{
	return ctx.JSON(200,echo.Map{
		"message":"success view product by id",
		"detail":"-",
	})
}

// Update by id
func (c *ProductController) UpdateProductById(ctx echo.Context) error{
	return ctx.JSON(200,echo.Map{
		"message":"success Update product by id",
		"detail":"-",
	})
}

// Delete by id
func (c *ProductController) DeleteProductById(ctx echo.Context) error{
	return ctx.JSON(200,echo.Map{
		"message":"success Delete product by id",
		"detail":"-",
	})
}
