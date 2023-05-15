package models

import (
	"gorm.io/gorm"
)

type ProductsSuppliers struct {
	ProductID  uint32    `gorm:"not null;primaryKey;"`
	SupplierID uint32    `gorm:"not null;primaryKey;"`
	Stock      uint32    `gorm:"column:stock;not null" json:"stock"`
	Products   Products  `gorm:"foreignKey:ProductID; constraint:OnDelete: CASCADE"`
	Suppliers  Suppliers `gorm:"foreignKey:SupplierID; constraint:OnDelete: CASCADE"`
}

type StockRequest struct {
	ProductID  uint32 `json:"product_id" binding:"required"`
	SupplierID uint32 `json:"supplier_id" binding:"required"`
	Stock      uint32 `json:"stock" binding:"required"`
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
	return db.Model(&ProductsSuppliers{}).Create(Stock).Error
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

func SelectSupplierStocks(db *gorm.DB, id int) ([]ProductsSuppliers, error) {
	var supplierStocks []ProductsSuppliers
	err := db.Where("supplier_id = ?", id).Find(&supplierStocks).Error
	if err != nil {
		return nil, err
	}

	return supplierStocks, nil
}

func GetProductsBySupplierID(db *gorm.DB, id uint32) ([]map[string]interface{}, error) {
	var products []map[string]interface{}
	err := db.Model(&Products{}).Select("products.id, products.name").
		Joins("left join products_suppliers on products_suppliers.product_id = products.id").
		Where("products_suppliers.supplier_id = ?", id).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetMissingProductsBySupplierID(db *gorm.DB, id uint32) ([]map[string]interface{}, error) {
	var products []map[string]interface{}
	var err error
	products, err = GetProductsBySupplierID(db, id)
	var productIDs []uint32
	for _, p := range products {
		productIDs = append(productIDs, p["id"].(uint32))
	}
	var missingProducts []map[string]interface{}
	err = db.Model(&Products{}).Select("id", "name").Not("id IN (?)", productIDs).Find(&missingProducts).Error
	if err != nil {
		return make([]map[string]interface{}, 0), err
	}
	return missingProducts, nil
}
