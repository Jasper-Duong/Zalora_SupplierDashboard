package validations

import (
	"fmt"
	"net/http"
	"server/internal/models"

	"github.com/gin-gonic/gin"
)

func AddressValidation(c *gin.Context) {
	var data models.AddressCreate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(data)
	c.Set("data", data)
	c.Next()

}
