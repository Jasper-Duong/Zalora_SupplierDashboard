package handlers

import (
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type ProductsHandler struct {
	BaseHandler
}

func NewProductsHandler() *ProductsHandler {
	return &ProductsHandler{}
}

func (h *ProductsHandler) GetProducts(c *gin.Context) {
	var query models.ProductsQueryParams
	err := c.ShouldBindQuery(&query)
	if err != nil {
		h.handleError(c, err)
		return
	}
	products, total, err := services.GetProducts(&query)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessGet(c, map[string]interface{}{
		"products": products,
		"total":    total,
	})
}

func (h *ProductsHandler) GetProductSuppliers(c *gin.Context) {
	suppliers, err := services.GetSuppliersByProductID(c.Param("product_id"))
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessGet(c, suppliers)
}

func (h *ProductsHandler) GetProductByID(c *gin.Context) {
	id := h.parseId(c, c.Param("product_id"))
	if id == 0 {
		return
	}
	product, err := services.GetProductByID(id)
	if err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessGet(c, &product)
}

func (h *ProductsHandler) CreateProduct(c *gin.Context) {
	var product models.Products
	if err := h.validateInput(c, &product); err != nil {
		return
	}

	if err := services.CreateProduct(&product); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessCreate(c)
}

func (h *ProductsHandler) UpdateProduct(c *gin.Context) {
	var product models.Products
	if err := h.validateInput(c, &product); err != nil {
		return
	}

	id := h.parseId(c, c.Param("product_id"))
	if id == 0 {
		return
	}

	if err := services.UpdateProduct(&product, id); err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessUpdate(c)
}

func (h *ProductsHandler) DeleteProduct(c *gin.Context) {
	id := h.parseId(c, c.Param("product_id"))
	if id == 0 {
		return
	}
	if err := services.DeleteProduct(id); err != nil {
		h.handleError(c, err)
		return
	}
	h.handleSuccessDelete(c)
}

func (h *ProductsHandler) GetAttributeOfProducts(c *gin.Context) {
	attribute := c.Param("attribute")
	products, err := services.GetAttributeOfProducts(attribute)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessGet(c, products)
}

func (h *ProductsHandler) GetProductStocks(c *gin.Context) {
	id := h.parseId(c, c.Param("product_id"))
	if id == 0 {
		return
	}
	stocks, err := services.GetProductStocks(id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessGet(c, stocks)
}

func (h *ProductsHandler) GetProductMissingSuppliers(c *gin.Context) {
	id := h.parseId(c, c.Param("product_id"))
	if id == 0 {
		return
	}
	suppliers, err := services.GetProductMissingSuppliers(id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	h.handleSuccessGet(c, suppliers)
}
