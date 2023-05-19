package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadProductsRoute(route *gin.RouterGroup) {
	productsHandler := handlers.NewProductsHandler()
	route.GET("/", productsHandler.GetProducts)
	route.GET("/:product_id", productsHandler.GetProductByID)
	route.POST("/", productsHandler.CreateProduct)
	route.PUT("/:product_id", productsHandler.UpdateProduct)
	route.DELETE("/:product_id", productsHandler.DeleteProduct)

	route.GET("/attribute/:attribute", productsHandler.GetAttributeOfProducts)

	curGroup := route.Group("/:product_id")
	{
		curGroup.GET("/stocks", productsHandler.GetProductStocks)
		curGroup.GET("/suppliers", productsHandler.GetProductSuppliers)
		curGroup.GET("/suppliers/missing", productsHandler.GetProductMissingSuppliers)
	}
}
