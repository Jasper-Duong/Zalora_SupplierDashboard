package handlers

import (
	"net/http"
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func GetSuppliersHandler(c *gin.Context) {
	var query models.QueryParam
	c.ShouldBindQuery(&query)
	suppliers, total, err := services.GetSuppliers(&query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := gin.H{
		"suppliers": suppliers,
		"total":     total,
	}

	c.JSON(http.StatusOK, res)
}

func CreateSuppliersHandler(c *gin.Context) {
	supplier := c.MustGet("supplier").(models.Suppliers)

	if status, err := services.CreateSupplier(&supplier); err != nil {
		c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supplier)
}

func UpdateSuppliers(c *gin.Context) {
	supplier := c.MustGet("supplier").(models.Suppliers)
	id := c.Param("id")

	if status, err := services.UpdateSupplier(&supplier, id); err != nil {
		c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "okei"})
}

func DeleteSuppliers(c *gin.Context) {
	/*supplier := c.MustGet("supplier").(models.Suppliers)
	id := c.Param("id")*/
}
