package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadSuppliersRoute(route *gin.RouterGroup) {
	route.GET("/", handlers.GetSuppliers)
}
