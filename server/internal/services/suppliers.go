package services

import (
	"fmt"
	"server/db"
	"server/internal/models"
)

func GetSuppliers(query *models.QueryParam) ([]models.Suppliers, error) {
	return models.SelectSuppliers(db.DB, query)
}

func CreateSupplier(supplier *models.Suppliers) error {
	fmt.Println(supplier)
	return models.CreateSupplier(db.DB, supplier)
}
