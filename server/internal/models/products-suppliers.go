package models

import "gorm.io/gorm"

// foreign key: product_id
type ProductsSuppliers struct {
	ProductsID  uint64 `gorm:"not null"`
	SuppliersID uint64 `gorm:"not null"`
	Stock       uint32 `gorm:"column:stock;not null" json:"stock"`
	Products    Products
	Suppliers   Suppliers
}

func SelectProductStocks(db *gorm.DB, id int) ([]ProductsSuppliers, error) {
	var productStocks []ProductsSuppliers
	err := db.Where("product_id = ?", id).Find(&productStocks).Error
	if err != nil {
		return nil, err
	}

	return productStocks, nil
}
