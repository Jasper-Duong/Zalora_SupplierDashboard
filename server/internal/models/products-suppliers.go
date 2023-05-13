package models

import (
	"gorm.io/gorm"
)

type ProductsSuppliers struct {
	ProductID  uint32 `gorm:"column:product_id;primaryKey" json:"product_id" binding:"required"`
	SupplierID uint32 `gorm:"column:supplier_id;primaryKey" json:"supplier_id" binding:"required"`
	Stock      uint32 `gorm:"column:stock;not null" json:"stock" binding:"required" binding:"required"`
}

func (current *ProductsSuppliers) AfterCreate(tx *gorm.DB) (err error) {
	productID := current.ProductID
	supplierID := current.SupplierID
	stock := current.Stock

	err = tx.Model(&Products{}).Where("id = ?", productID).Update("stock", gorm.Expr("stock + ?", stock)).Error
	if err != nil {
		return err
	}

	err = tx.Model(&Suppliers{}).Where("id = ?", supplierID).Update("stock", gorm.Expr("stock + ?", stock)).Error
	if err != nil {
		return err
	}

	return nil
}

func (old *ProductsSuppliers) BeforeUpdate(tx *gorm.DB) (err error) {
	productID := old.ProductID
	supplierID := old.SupplierID
	stock := old.Stock

	err = tx.Model(&Products{}).Where("id = ?", productID).Update("stock", gorm.Expr("stock - ?", stock)).Error
	if err != nil {
		return err
	}

	err = tx.Model(&Suppliers{}).Where("id = ?", supplierID).Update("stock", gorm.Expr("stock - ?", stock)).Error
	if err != nil {
		return err
	}

	return nil
}

func (new *ProductsSuppliers) AfterUpdate(tx *gorm.DB) (err error) {
	productID := new.ProductID
	supplierID := new.SupplierID
	stock := new.Stock

	err = tx.Model(&Products{}).Where("id = ?", productID).Update("stock", gorm.Expr("stock + ?", stock)).Error
	if err != nil {
		return err
	}

	err = tx.Model(&Suppliers{}).Where("id = ?", supplierID).Update("stock", gorm.Expr("stock + ?", stock)).Error
	if err != nil {
		return err
	}

	return nil
}

func (current *ProductsSuppliers) AfterDelete(tx *gorm.DB) (err error) {
	productID := current.ProductID
	supplierID := current.SupplierID
	stock := current.Stock

	err = tx.Model(&Products{}).Where("id = ?", productID).Update("stock", gorm.Expr("stock - ?", stock)).Error
	if err != nil {
		return err
	}

	err = tx.Model(&Suppliers{}).Where("id = ?", supplierID).Update("stock", gorm.Expr("stock - ?", stock)).Error
	if err != nil {
		return err
	}

	return nil
}

func SelectProductStocks(db *gorm.DB, id int) ([]ProductsSuppliers, error) {
	var stocks []ProductsSuppliers
	err := db.Where("product_id = ?", id).Find(&stocks).Error
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func CreateStock(db *gorm.DB, Stock *ProductsSuppliers) error {
	return db.Create(Stock).Error
}

func UpdateStock(db *gorm.DB, product_id uint32, supplier_id uint32, stock uint32) error {
	old := ProductsSuppliers{}
	err := db.Where("product_id = ? AND supplier_id = ?", product_id, supplier_id).First(&old).Error
	if err != nil {
		return err
	}

	new := ProductsSuppliers{
		ProductID:  product_id,
		SupplierID: supplier_id,
		Stock:      stock,
	}

	return db.Model(&old).Updates(new).Error
}

func DeleteStock(db *gorm.DB, product_id uint32, supplier_id uint32) error {
	current := ProductsSuppliers{}
	err := db.Where("product_id = ? AND supplier_id = ?", product_id, supplier_id).First(&current).Error
	if err != nil {
		return err
	}
	return db.Where("product_id = ? AND supplier_id = ?", product_id, supplier_id).Delete(&current).Error
}

func DeleteStockByProductID(db *gorm.DB, product_id uint32) error {
	stocks := []ProductsSuppliers{}
	err := db.Where("product_id = ?", product_id).Find(&stocks).Error
	if err != nil {
		return err
	}
	for _, stock := range stocks {
		err = db.Where("product_id = ? AND supplier_id = ?", stock.ProductID, stock.SupplierID).Delete(&stock).Error
		if err != nil {
			return err
		}
	}
	return nil
}
