package services

import (
	"fmt"
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
	supplier.ID = uint32(ID)
	return models.UpdateSupplier(db.DB, supplier)
}

func DeleteSupplier(id string) error {
	ID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	return models.DeleteSupplier(db.DB, ID)
}

func GetSuppliersName() ([]map[string]interface{}, error) {
	return models.GetSuppliersAttribute(db.DB, "name")
}

func GetSupplierAddresses(id string) ([]models.Addresses, error) {
	ID_, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return []models.Addresses{}, err
	}
	ID := uint32(ID_)
	if err = models.GetSupplierByID(db.DB, ID); err != nil {
		return []models.Addresses{}, err
	}
	return models.GetAddressesBySupplierID(db.DB, ID)
}

func GetSupplierMissingProducts(id string) ([]map[string]interface{}, error) {
	ID_, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return make([]map[string]interface{}, 0), err
	}
	ID := uint32(ID_)
	var products []map[string]interface{}
	if products, err = models.GetMissingProductsBySupplierID(db.DB, ID); err != nil {
		fmt.Println(products)
	}
	return products, nil
}
