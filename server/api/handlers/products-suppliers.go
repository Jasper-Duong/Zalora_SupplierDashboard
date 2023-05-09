package handlers

import (
	"net/http"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func GetProductStocks(c *gin.Context) {
	id := c.Param("id")
	stocks, err := services.GetProductStocks(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stocks)
}
