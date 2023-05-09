package models

import "gorm.io/gorm"

// foreign key: product_id
type ProductsSuppliers struct {
	ProductID  uint64 `gorm:"column:product_id;primaryKey" json:"product_id"`
	SupplierID uint64 `gorm:"column:supplier_id;primaryKey" json:"supplier_id"`
	Stock      uint32 `gorm:"column:stock;not null" json:"stock"`
}

func SelectProductStocks(db *gorm.DB, id int) ([]ProductsSuppliers, error) {
	var productStocks []ProductsSuppliers
	err := db.Where("product_id = ?", id).Find(&productStocks).Error
	if err != nil {
		return nil, err
	}

	return productStocks, nil
}
