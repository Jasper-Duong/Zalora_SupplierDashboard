package handlers

import (
	"fmt"
	"net/http"
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func GetSuppliersHandler(c *gin.Context) {
	var query models.QueryParam
	c.ShouldBindQuery(&query)
	fmt.Println(query)
	suppliers, err := services.GetSuppliers(&query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, suppliers)
}

func CreateSuppliersHandler(c *gin.Context) {
	var supplier models.Suppliers

	if err := c.BindJSON(&supplier); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateSupplier(&supplier); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"supplier": supplier})
}

func UpdateSuppliers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "okei",
	})
}
