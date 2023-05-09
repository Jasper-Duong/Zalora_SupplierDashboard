package routes

import (
	"server/api/handlers"
	"server/api/middlewares/validations"

	"github.com/gin-gonic/gin"
)

func LoadSuppliersRoute(route *gin.RouterGroup) {
	route.GET("/", handlers.GetSuppliersHandler)
	route.POST("/", validations.SupplierValidation, handlers.CreateSuppliersHandler)
	route.PATCH("/:id", validations.SupplierValidation, handlers.UpdateSuppliers)
	route.DELETE("/:id", handlers.DeleteSuppliers)
}
