package services

import (
	"fmt"
	"server/internal/models"
	"testing"
)

func TestCreateSupplier(t *testing.T) {
	arg := &models.Suppliers{
		Name:          "help",
		Email:         "HELP",
		ContactNumber: "0786948444",
	}
	err := models.CreateSupplier(DBtest, arg)
	fmt.Print(err)
}
