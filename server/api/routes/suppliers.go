package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadSuppliersRoute(route *gin.RouterGroup) {
	route.GET("/", handlers.GetSuppliersHandler)
	route.POST("/", handlers.CreateSuppliersHandler)
	route.PUT("/:id", handlers.UpdateSuppliers)
	route.DELETE("/:id", handlers.DeleteSupplier)

	route.GET("/attribute/name", handlers.GetSuppliersName)
	route.GET("/:id/stocks", handlers.GetSupplierStocks)
	route.GET("/:id/addresses", handlers.GetSupplierAddresses)
	route.GET("/:id/products/missing", handlers.GetSupplierMissingProducts)
}
