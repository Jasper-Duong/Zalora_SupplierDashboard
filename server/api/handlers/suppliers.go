package handlers

import (
	"errors"
	"net/http"
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetSuppliersHandler(c *gin.Context) {
	var query models.SuppliersQueryParam
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
	var supplier models.Suppliers
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateSupplier(&supplier); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func UpdateSuppliers(c *gin.Context) {
	var supplier models.Suppliers
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")

	if err := services.UpdateSupplier(&supplier, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteSupplier(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteSupplier(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func GetSuppliersName(c *gin.Context) {
	names, err := services.GetSuppliersName()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, names)
}

func GetSupplierAddresses(c *gin.Context) {
	id := c.Param("id")
	addresses, err := services.GetSupplierAddresses(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, addresses)
}

func GetSupplierMissingProducts(c *gin.Context) {
	id := c.Param("id")
	products, _ := services.GetSupplierMissingProducts(id)
	c.JSON(http.StatusOK, products)
}
