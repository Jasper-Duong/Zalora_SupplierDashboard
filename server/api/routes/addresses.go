package routes

import (
	"server/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadAddressesRoute(route *gin.RouterGroup) {
	addressesHandler := handlers.NewAddressesHandler()
	route.GET("/", addressesHandler.GetAddresses)
	route.POST("/", addressesHandler.CreateAddress)
	route.PUT("/:id", addressesHandler.UpdateAddress)
	route.DELETE("/:id", addressesHandler.DeleteAddress)
}
