package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadProductsRoute(route *gin.RouterGroup) {
	route.GET("/", handlers.GetProducts)
}
