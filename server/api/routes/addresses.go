package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadAddressesRoute(route *gin.RouterGroup) {
	route.GET("/", handlers.GetAddresses)
	route.POST("/", handlers.CreateAddress)
	route.PUT("/:id", handlers.UpdateAddress)
	route.DELETE("/:id", handlers.DeleteAddress)
}
