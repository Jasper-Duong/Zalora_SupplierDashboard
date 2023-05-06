package handlers

import "github.com/gin-gonic/gin"

func GetSuppliers(c *gin.Context) {
	c.JSON(202, gin.H{
		"message": "okei",
	})
}
