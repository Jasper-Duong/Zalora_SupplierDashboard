package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadProductsRoute(route *gin.RouterGroup) {
	route.GET("/", handlers.GetProducts)
	route.POST("/", handlers.CreateProduct)
	route.PUT("/:id", handlers.UpdateProduct)
	route.DELETE("/:id", handlers.DeleteProduct)

	route.GET("/:id/suppliers", handlers.GetProductStocks)

}
