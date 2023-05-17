package handlers

import (
	"errors"
	"net/http"

	"server/internal/models"
	"server/internal/services"
	"server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductListResponse struct {
	TotalCount uint32                         `json:"total"`
	Products   []models.ProductsWithSuppliers `json:"products"`
}

func GetProducts(c *gin.Context) {
	var query models.QueryParams
	c.ShouldBindQuery(&query)
	products, total, err := services.GetProducts(&query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := ProductListResponse{
		TotalCount: total,
		Products:   products,
	}
	c.JSON(http.StatusOK, response)
}

func GetProductSuppliers(c *gin.Context) {
	suppliers, err := services.GetSuppliersByProductID(c.Param("product_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, suppliers)
}

func GetProductByID(c *gin.Context) {
	product, err := services.GetProductByID(c.Param("product_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Products
	if err := utils.ValidateInput(&product, c); err != nil {
		return
	}

	if err := services.CreateProduct(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func UpdateProduct(c *gin.Context) {
	var product models.Products
	if err := utils.ValidateInput(&product, c); err != nil {
		return
	}

	if err := services.UpdateProduct(&product, c.Param("product_id")); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func DeleteProduct(c *gin.Context) {
	if err := services.DeleteProduct(c.Param("product_id")); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func GetAttributeOfProducts(c *gin.Context) {
	attribute := c.Param("attribute")
	products, err := services.GetAttributeOfProducts(attribute)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
