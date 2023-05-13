package handlers

import (
	"errors"
	"net/http"
	"server/internal/models"
	"server/internal/services"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAddresses(c *gin.Context) {
	address := models.GetAddresses()
	c.JSON(http.StatusOK, address)
}

func CreateAddress(c *gin.Context) {
	var address models.AddressCreate
	// Bind req.Body -> address & validate required fields
	if err := utils.ValidateInput(&address, c); err != nil {
		return
	}

	// Call service -> CreateAddress
	if err := services.CreateAddress(&address); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Return response with statusCreated & data
	c.Status(http.StatusCreated)
}

func UpdateAddress(c *gin.Context) {
	var address models.AddressUpdate
	if err := utils.ValidateInput(&address, c); err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := services.UpdateAddress(&address, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteAddress(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := services.DeleteAddress(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)

}
