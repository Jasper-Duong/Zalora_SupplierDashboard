package services

import (
	"server/db"
	"server/internal/models"
	"strconv"
)

func GetSuppliers(query *models.QueryParam) ([]models.Suppliers, int64, error) {
	return models.SelectSuppliers(db.DB, query)
}

func CreateSupplier(supplier *models.Suppliers) (int, error) {
	return models.CreateSupplier(db.DB, supplier)
}

func UpdateSupplier(supplier *models.Suppliers, id string) (int, error) {
	ID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 500, err
	}
	supplier.ID = uint64(ID)
	return models.UpdateSupplier(db.DB, supplier)
}

func DeleteSupplier(id string) (int, error) {
	ID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 500, err
	}
	return models.DeleteSupplier(db.DB, ID)
}

func GetSuppliersName() ([]models.Suppliers, error) {
	return models.GetSuppliersName(db.DB)
}
