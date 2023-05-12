package models

import "gorm.io/gorm"

// foreign key: product_id
type ProductsSuppliers struct {
	ProductID  uint64    `gorm:"not null;primaryKey;"`
	SupplierID uint64    `gorm:"not null;primaryKey;"`
	Stock      uint32    `gorm:"column:stock;not null" json:"stock"`
	Products   Products  `gorm:"foreignKey:ProductID"`
	Suppliers  Suppliers `gorm:"foreignKey:SupplierID"`
}

func SelectProductStocks(db *gorm.DB, id int) ([]ProductsSuppliers, error) {
	var productStocks []ProductsSuppliers
	err := db.Where("product_id = ?", id).Find(&productStocks).Error
	if err != nil {
		return nil, err
	}

	return productStocks, nil
}

func SelectSupplierStocks(db *gorm.DB, id int) ([]ProductsSuppliers, error) {
	var supplierStocks []ProductsSuppliers
	err := db.Where("supplier_id = ?", id).Find(&supplierStocks).Error
	if err != nil {
		return nil, err
	}

	return supplierStocks, nil
}
