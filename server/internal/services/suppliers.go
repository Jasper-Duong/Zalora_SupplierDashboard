package services

import (
	"server/db"
	"server/internal/models"
	"strconv"
)

func GetSuppliers(query *models.SuppliersQueryParam) ([]models.Suppliers, int64, error) {
	return models.GetSuppliers(db.DB, query)
}

func CreateSupplier(supplier *models.Suppliers) error {
	return models.CreateSupplier(db.DB, supplier)
}

func UpdateSupplier(supplier *models.Suppliers, id string) error {
	ID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	supplier.ID = uint64(ID)
	return models.UpdateSupplier(db.DB, supplier)
}

func DeleteSupplier(id string) error {
	ID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	return models.DeleteSupplier(db.DB, ID)
}

func GetSuppliersName() ([]models.SuppliersInfo, error) {
	return models.GetSuppliersName(db.DB)
}
