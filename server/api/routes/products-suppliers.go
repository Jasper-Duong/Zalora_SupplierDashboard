package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadProductsSuppliersRoute(route *gin.RouterGroup) {
	route.POST("/", handlers.CreateStock)
	route.PUT("/", handlers.UpdateStock)
	curGroup := route.Group("/:product_id")
	{
		curGroup.DELETE("/:supplier_id", handlers.DeleteStock)
	}

}
