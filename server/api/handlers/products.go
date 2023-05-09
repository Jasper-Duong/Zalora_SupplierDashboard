package handlers

import (
	"net/http"

	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type ProductListResponse struct {
	TotalCount int64             `json:"total"`
	Products   []models.Products `json:"products"`
}

func GetProducts(c *gin.Context) {
	var query models.QueryParams
	c.ShouldBindQuery(&query)
	products, total, err := services.GetProducts(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := ProductListResponse{
		TotalCount: total,
		Products:   products,
	}
	c.JSON(http.StatusOK, response)
}

func CreateProduct(c *gin.Context) {
	var product models.Products
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.CreateProduct(&product)

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func UpdateProduct(c *gin.Context) {
	var product models.Products

	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.UpdateProduct(&product, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProduct(c *gin.Context) {
	err := services.DeleteProduct(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
