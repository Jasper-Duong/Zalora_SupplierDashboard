package tests

import (
	"server/api/routes"
	"testing"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func TestMain(t *testing.T) {
	ConnectDB()
	r = gin.Default()
	routes.LoadRoutes(r)
}
