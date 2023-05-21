package services

import (
	"server/db"
	"server/internal/models"
)

var GetSuppliers = func(query *models.SuppliersQueryParam) ([]models.SupplierWithAddresses, uint32, error) {
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

var CreateSupplier = func(supplier *models.Suppliers) error {
	return models.CreateSupplier(db.DB, supplier)
}

var UpdateSupplier = func(supplier *models.Suppliers, id uint32) error {
	_, err := models.GetSupplierByID(db.DB, id)
	if err != nil {
		return err
	}
	supplier.ID = id
	return models.UpdateSupplier(db.DB, supplier)
}

var DeleteSupplier = func(id uint32) error {
	return models.DeleteSupplier(db.DB, id)
}

var GetSupplierByID = func(id uint32) (models.Suppliers, error) {
	return models.GetSupplierByID(db.DB, id)
}

var GetSuppliersName = func() ([]map[string]interface{}, error) {
	return models.GetSuppliersAttribute(db.DB, "name")
}

var GetSupplierAddresses = func(id uint32) ([]map[string]interface{}, error) {
	if _, err := models.GetSupplierByID(db.DB, id); err != nil {
		return make([]map[string]interface{}, 0), err
	}
	return models.GetAddressesBySupplierID(db.DB, id)
}

/*func GetSupplierMissingProducts(id uint32) ([]map[string]interface{}, error) {
	var products []map[string]interface{}
	if products, err := models.GetSupplierMissingProducts(db.DB, id); err != nil {
		fmt.Println(products)
	}
	return products, nil
}*/
