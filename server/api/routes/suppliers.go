package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadSuppliersRoute(route *gin.RouterGroup) {
	route.GET("/", handlers.GetSuppliersHandler)
	route.POST("/", handlers.CreateSuppliersHandler)
	route.PATCH("/:id", handlers.UpdateSuppliers)
	route.DELETE("/:id", handlers.DeleteSupplier)

	route.GET("/attribute/name", handlers.GetSuppliersName)
	route.GET("/:id/stocks", handlers.GetSupplierStocks)
}
