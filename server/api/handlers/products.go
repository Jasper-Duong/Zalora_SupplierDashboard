package handlers

import (
	"net/http"
	"strconv"

	"server/internal/services"

	"server/internal/models"

	"github.com/gin-gonic/gin"
)

type ProductListResponse struct {
	TotalCount int64             `json:"total"`
	Products   []models.Products `json:"products"`
}

func GetProducts(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		offset = 0
	}

	params := c.Request.URL.Query()
	delete(params, "limit")
	delete(params, "offset")

	products, total, err := services.GetProducts(limit, offset, params)

	response := ProductListResponse{
		TotalCount: total,
		Products:   products,
	}

	c.JSON(http.StatusOK, response)
}
