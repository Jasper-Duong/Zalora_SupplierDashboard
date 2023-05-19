package handlers

import (
	"net/http"
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

type AddressesHandler struct {
	BaseHandler
}

func NewAddressesHandler() *AddressesHandler {
	return &AddressesHandler{}
}

func (h *AddressesHandler) GetAddresses(c *gin.Context) {
	address := models.GetAddresses()
	c.JSON(http.StatusOK, address)
}

func (h *AddressesHandler) CreateAddress(c *gin.Context) {
	var address models.AddressCreate
	// Bind req.Body -> address & validate required fields
	if err := h.validateInput(c, &address); err != nil {
		return
	}

	// Call service -> CreateAddress
	if err := services.CreateAddress(&address); err != nil {
		h.handleError(c, err)
	}

	// Return response with statusCreated & data
	h.handleSuccessCreate(c)
}

func (h *AddressesHandler) UpdateAddress(c *gin.Context) {
	var address models.AddressUpdate
	if err := h.validateInput(c, &address); err != nil {
		return
	}

	id := h.parseId(c, c.Param("id"))
	if id == 0 {
		return
	}

	if err := services.UpdateAddress(&address, id); err != nil {
		h.handleError(c, err)
	}
	h.handleSuccessCreate(c)
}

func (h *AddressesHandler) DeleteAddress(c *gin.Context) {
	id := h.parseId(c, c.Param("id"))
	if id == 0 {
		return
	}

	if err := services.DeleteAddress(id); err != nil {
		h.handleError(c, err)
	}
	c.Status(http.StatusOK)

}
