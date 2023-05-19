package handlers

import (
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type ProductsSuppliersHandler struct {
	BaseHandler
}

func NewProductsSuppliersHandler() *ProductsSuppliersHandler {
	return &ProductsSuppliersHandler{}
}

func (h *ProductsSuppliersHandler) CreateStock(c *gin.Context) {
	var stock models.StockRequest
	if err := h.validateInput(c, &stock); err != nil {
		return
	}

	if err := services.CreateStock(stock); err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessCreate(c)
}

func (h *ProductsSuppliersHandler) UpdateStock(c *gin.Context) {
	var stock models.StockRequest

	if err := h.validateInput(c, &stock); err != nil {
		return
	}

	if err := services.UpdateStock(stock); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessUpdate(c)
}

func (h *ProductsSuppliersHandler) DeleteStock(c *gin.Context) {
	product_id := h.parseId(c, c.Param("product_id"))
	if product_id == 0 {
		return
	}
	supplier_id := h.parseId(c, c.Param("supplier_id"))
	if supplier_id == 0 {
		return
	}

	if err := services.DeleteStock(product_id, supplier_id); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessDelete(c)
}
