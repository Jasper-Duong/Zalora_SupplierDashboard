package models

import "gorm.io/gorm"

type Products struct {
	ID     uint32 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name   string `gorm:"column:name;not null" json:"name"`
	Brand  string `gorm:"column:brand;not null" json:"brand"`
	Sku    string `gorm:"column:SKU;not null" json:"sku"`
	Size   string `gorm:"column:size;not null" json:"size"`
	Color  string `gorm:"column:color;not null" json:"color"`
	Status bool   `gorm:"column:status;not null" json:"status"`
	Stock  uint32 `gorm:"column:stock;default:0" json:"stock"`
}

type QueryParams struct {
	Page   int      `form:"page"`
	Limit  int      `form:"limit"`
	Brand  []string `form:"brand[]"`
	Size   []string `form:"size[]"`
	Color  []string `form:"color[]"`
	Status []string `form:"status[]"`
}

func SelectProducts(db *gorm.DB, query *QueryParams) ([]Products, int64, error) {
	var total int64
	err := db.Model(&Products{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var products []Products
	offset := (query.Page - 1) * query.Limit
	db = db.Limit(query.Limit).Offset(offset)

	err = db.Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func CreateProduct(db *gorm.DB, product *Products) error {
	return db.Create(product).Error
}

func UpdateProduct(db *gorm.DB, product *Products, id int) error {
	return db.Model(&product).Where("id = ?", id).Updates(product).Error
}

func DeleteProduct(db *gorm.DB, id int) error {
	return db.Delete(&Products{}, id).Error
}
