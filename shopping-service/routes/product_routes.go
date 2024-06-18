package routes

import (
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/controller"
	"github.com/labstack/echo/v4"
)

func ProductRoute(e *echo.Echo, productController controller.ProductController) {
	routes := e.Group("/api")

	routesV1 := routes.Group("/v1")
	// Create New Product
	routesV1.POST("/products", productController.HandleCreateProductRequest)
	// Get All Porduct with pagging
	routesV1.GET("/all-products", productController.HandleGetProductsWithPagingRequest)
	// Get product by ID
	routesV1.GET("/products/:id", productController.HandleGetProductByIDRequest)
	// Update Product by ID
	routesV1.PUT("/products/:id", productController.HandleUpdateProductRequest)
	// Delete product by ID
	routesV1.DELETE("/products/:id", productController.HandleDeleteProductRequest)
}
