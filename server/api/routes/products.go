package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadProductsRoute(route *gin.RouterGroup) {
	route.GET("/", handlers.GetProducts)
	route.GET("/:product_id", handlers.GetProductByID)
	route.POST("/", handlers.CreateProduct)
	route.PUT("/:product_id", handlers.UpdateProduct)
	route.DELETE("/:product_id", handlers.DeleteProduct)

	route.GET("/attribute/:attribute", handlers.GetAttributeOfProducts)

	curGroup := route.Group("/:product_id")
	{
		curGroup.GET("/suppliers", handlers.GetProductStocks)
		curGroup.POST("/suppliers", handlers.CreateProductStock)
		curGroup.PUT("/suppliers/:supplier_id", handlers.UpdateProductStock)
		curGroup.DELETE("/suppliers/:supplier_id", handlers.DeleteProductStock)
	}
}