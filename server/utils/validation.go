package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateInput[T any](input *T, c *gin.Context) error {
	if err := c.ShouldBindJSON(input); err != nil {
		fmt.Println(input)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return err
	}
	return nil
}
