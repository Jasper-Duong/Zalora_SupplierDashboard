package handlers

import (
	"net/http"
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func GetProductStocks(c *gin.Context) {
	id := c.Param("product_id")
	stocks, err := services.GetProductStocks(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stocks)
}

func CreateStock(c *gin.Context) {
	var stock models.ProductsSuppliers
	err := c.ShouldBindJSON(&stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = services.CreateStock(stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock created"})
}

func UpdateStock(c *gin.Context) {
	var stock models.ProductsSuppliers
	err := c.ShouldBindJSON(&stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = services.UpdateStock(stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock updated"})
}

func DeleteStock(c *gin.Context) {
	product_id := c.Param("product_id")
	supplier_id := c.Param("supplier_id")

	err := services.DeleteStock(product_id, supplier_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock deleted"})
}
