package validations

import (
	"net/http"
	"server/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SupplierValidation(c *gin.Context) {
	var supplier models.Suppliers
	if err := c.BindJSON(&supplier); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(&supplier); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Set("supplier", supplier)
	c.Next()
}
