package handlers

import (
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type SuppliersHandler struct {
	BaseHandler
}

func NewSuppliersHandler() *SuppliersHandler {
	return &SuppliersHandler{}
}

func (h *SuppliersHandler) GetSuppliers(c *gin.Context) {
	var query models.SuppliersQueryParam
	c.ShouldBindQuery(&query)
	suppliers, total, err := services.GetSuppliers(&query)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessGet(c, map[string]interface{}{
		"suppliers": suppliers,
		"total":     total,
	})
}

func (h *SuppliersHandler) CreateSupplier(c *gin.Context) {
	var supplier models.Suppliers
	if err := h.validateInput(c, &supplier); err != nil {
		return
	}

	if err := services.CreateSupplier(&supplier); err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessCreate(c)
}

func (h *SuppliersHandler) UpdateSupplier(c *gin.Context) {
	var supplier models.Suppliers
	if err := h.validateInput(c, &supplier); err != nil {
		return
	}
	id := h.parseId(c, c.Param("id"))
	if id == 0 {
		return
	}
	if err := services.UpdateSupplier(&supplier, id); err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessUpdate(c)
}

func (h *SuppliersHandler) DeleteSupplier(c *gin.Context) {
	id := h.parseId(c, c.Param("id"))
	if id == 0 {
		return
	}

	if err := services.DeleteSupplier(id); err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessDelete(c)
}

func (h *SuppliersHandler) GetSuppliersName(c *gin.Context) {
	names, err := services.GetSuppliersName()
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessGet(c, &names)
}

func (h *SuppliersHandler) GetSupplierByID(c *gin.Context) {
	id := h.parseId(c, c.Param("id"))
	if id == 0 {
		return
	}
	supplier, err := services.GetSupplierByID(id)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessGet(c, &supplier)
}

func (h *SuppliersHandler) GetSupplierStocks(c *gin.Context) {
	id := h.parseId(c, c.Param("id"))
	if id == 0 {
		return
	}
	stocks, err := services.GetSupplierStocks(id)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessGet(c, stocks)

}

func (h *SuppliersHandler) GetSupplierAddresses(c *gin.Context) {
	id := h.parseId(c, c.Param("id"))
	if id == 0 {
		return
	}
	addresses, err := services.GetSupplierAddresses(id)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessGet(c, addresses)
}

func (h *SuppliersHandler) GetSupplierMissingProducts(c *gin.Context) {
	id := h.parseId(c, c.Param("id"))
	if id == 0 {
		return
	}
	products, err := services.GetSupplierMissingProducts(id)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessGet(c, products)
}
