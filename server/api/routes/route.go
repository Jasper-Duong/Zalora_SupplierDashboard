package routes

import (
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	LoadSuppliersRoute(r.Group("/suppliers"))
	LoadAddressesRoute(r.Group("/addresses"))
}
