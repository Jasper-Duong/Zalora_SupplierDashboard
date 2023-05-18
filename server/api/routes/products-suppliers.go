package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadProductsSuppliersRoute(route *gin.RouterGroup) {
	productssuppliersHandler := handlers.NewProductsSuppliersHandler()
	route.POST("/", productssuppliersHandler.CreateStock)
	route.PUT("/", productssuppliersHandler.UpdateStock)
	curGroup := route.Group("/:product_id")
	{
		curGroup.DELETE("/:supplier_id", productssuppliersHandler.DeleteStock)
	}

}
