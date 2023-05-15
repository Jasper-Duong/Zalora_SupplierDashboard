package handlers

import (
	"net/http"
	"server/internal/models"
	"server/internal/services"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func GetProductStocks(c *gin.Context) {
	id := c.Param("product_id")
	stocks, err := services.GetProductStocks(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stocks)
}

func GetProductMissingSuppliers(c *gin.Context) {
	id := c.Param("product_id")
	suppliers, err := services.GetProductMissingSuppliers(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, suppliers)
}

func CreateStock(c *gin.Context) {
	var stock models.StockRequest
	if err := utils.ValidateInput(&stock, c); err != nil {
		return
	}

	if err := services.CreateStock(stock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func GetSupplierStocks(c *gin.Context) {
	id := c.Param("id")
	stocks, err := services.GetSupplierStocks(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stocks)

}

func UpdateStock(c *gin.Context) {
	var stock models.StockRequest

	if err := utils.ValidateInput(&stock, c); err != nil {
		return
	}

	if err := services.UpdateStock(stock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func DeleteStock(c *gin.Context) {
	product_id := c.Param("product_id")
	supplier_id := c.Param("supplier_id")

	if err := services.DeleteStock(product_id, supplier_id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
