package validations

import (
	"net/http"
	"server/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ProductValidation(c *gin.Context) {
	var product models.Products
	if err := c.BindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Set("product", product)
	c.Next()
}
