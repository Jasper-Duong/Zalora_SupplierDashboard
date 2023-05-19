package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadSuppliersRoute(route *gin.RouterGroup) {
	suppliersHandler := handlers.NewSuppliersHandler()
	route.GET("/", suppliersHandler.GetSuppliers)
	route.POST("/", suppliersHandler.CreateSupplier)

	route.GET("/:id", suppliersHandler.GetSupplierByID)
	route.PUT("/:id", suppliersHandler.UpdateSupplier)
	route.DELETE("/:id", suppliersHandler.DeleteSupplier)

	route.GET("/attribute/name", suppliersHandler.GetSuppliersName)
	route.GET("/:id/stocks", suppliersHandler.GetSupplierStocks)
	route.GET("/:id/addresses", suppliersHandler.GetSupplierAddresses)
	route.GET("/:id/products/missing", suppliersHandler.GetSupplierMissingProducts)
}
