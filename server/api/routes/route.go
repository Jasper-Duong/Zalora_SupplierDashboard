package routes

import (
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	superGroup := r.Group("/api")
	{
		LoadSuppliersRoute(superGroup.Group("/suppliers"))
	}
}
