package services

import (
	"fmt"
	"server/db"
	"server/internal/models"
)

func GetSuppliers(query *models.SuppliersQueryParam) ([]models.SupplierWithAddresses, uint32, error) {
	suppliers, total, err := models.GetSuppliers(db.DB, query)
	if err != nil {
		return nil, 0, err
	}
	var suppliersWithAddresses []models.SupplierWithAddresses
	for _, supplier := range suppliers {
		addresses, err := models.GetAddressesBySupplierID(db.DB, supplier.ID)
		if err != nil {
			return nil, 0, err
		}
		suppliersWithAddresses = append(suppliersWithAddresses, models.SupplierWithAddresses{
			Suppliers: supplier,
			Addresses: addresses,
		})
	}
	return suppliersWithAddresses, total, err
}

func CreateSupplier(supplier *models.Suppliers) error {
	return models.CreateSupplier(db.DB, supplier)
}

func UpdateSupplier(supplier *models.Suppliers, id uint32) error {
	_, err := models.GetSupplierByID(db.DB, id)
	if err != nil {
		return err
	}
	supplier.ID = id
	return models.UpdateSupplier(db.DB, supplier)
}

func DeleteSupplier(id uint32) error {
	return models.DeleteSupplier(db.DB, id)
}

func GetSupplierByID(id uint32) (models.Suppliers, error) {
	return models.GetSupplierByID(db.DB, id)
}

func GetSuppliersName() ([]map[string]interface{}, error) {
	return models.GetSuppliersAttribute(db.DB, "name")
}

func GetSupplierAddresses(id uint32) ([]map[string]interface{}, error) {
	if _, err := models.GetSupplierByID(db.DB, id); err != nil {
		return make([]map[string]interface{}, 0), err
	}
	return models.GetAddressesBySupplierID(db.DB, id)
}

func GetSupplierMissingProducts(id uint32) ([]map[string]interface{}, error) {
	var products []map[string]interface{}
	if products, err := models.GetMissingProductsBySupplierID(db.DB, id); err != nil {
		fmt.Println(products)
	}
	return products, nil
}
