package utils

import (
	"server/api/routes"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.LoadRoutes(r)
	return r
}
