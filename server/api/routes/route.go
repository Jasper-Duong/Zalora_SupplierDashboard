package routes

import (
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	superGroup := r.Group("/api")
	{
		LoadProductsRoute(superGroup.Group("/products"))
		LoadSuppliersRoute(superGroup.Group("/suppliers"))
	}
}
