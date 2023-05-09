package handlers

import (
	"net/http"
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

func CreateProductStock(c *gin.Context) {
	var stock map[string]interface{}
	err := c.ShouldBindJSON(&stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = services.CreateProductStock(stock, c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Stock created"})
}

func UpdateProductStock(c *gin.Context) {
	var data map[string]interface{}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = services.UpdateProductStock(data, c.Param("product_id"), c.Param("supplier_id"))
}

func DeleteProductStock(c *gin.Context) {
	product_id := c.Param("product_id")
	supplier_id := c.Param("supplier_id")

	err := services.DeleteProductStock(product_id, supplier_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock deleted"})
}
