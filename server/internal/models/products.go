package models

import (
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type Products struct {
	ID     uint32 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name   string `gorm:"column:name;not null" json:"name" binding:"required"`
	Brand  string `gorm:"column:brand;not null" json:"brand" binding:"required"`
	Sku    string `gorm:"column:sku;not null;unique" json:"sku"`
	Size   string `gorm:"column:size;not null" json:"size" binding:"required"`
	Color  string `gorm:"column:color;not null" json:"color" binding:"required"`
	Status *bool  `gorm:"column:status;not null" json:"status" binding:"required"`
	Stock  uint32 `gorm:"column:stock;default:0" json:"stock"`
}

type ProductsWithSuppliers struct {
	Products
	Suppliers []map[string]interface{} `json:"suppliers"`
}

type ProductsQueryParams struct {
	Page   int      `form:"page"`
	Limit  int      `form:"limit"`
	Status bool     `form:"status"`
	Name   string   `form:"name[]"`
	Brand  string   `form:"brand[]"`
	SKU    string   `form:"sku[]"`
	Size   []string `form:"size[]"`
	Color  []string `form:"color[]"`
}

func GetProducts(db *gorm.DB, query *ProductsQueryParams) ([]Products, uint32, error) {
	var products []Products
	db = db.Where("status = ?", query.Status)

	q := reflect.TypeOf(query).Elem()
	for i := 0; i < q.NumField(); i++ {
		field := q.Field(i)
		name := field.Name
		value := reflect.ValueOf(query).Elem().FieldByName(name)
		if name == "Page" || name == "Limit" || name == "Status" {
			continue
		}
		if value.Kind() == reflect.Slice && value.Len() > 0 {
			db = db.Where(strings.ToLower(name)+" IN ?", value.Interface())
		} else if value.Kind() != reflect.Slice && value.String() != "" {
			db = db.Where(strings.ToLower(name)+" LIKE ?", "%"+value.String()+"%")
		}
	}

	var total int64
	db.Model(&Products{}).Count(&total)
	offset := (query.Page - 1) * query.Limit
	db = db.Offset(offset).Limit(query.Limit)
	err := db.Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, uint32(total), nil
}

func GetProductsBySupplierID(db *gorm.DB, id uint32) ([]map[string]interface{}, error) {
	var products []map[string]interface{} = make([]map[string]interface{}, 0)
	err := db.Model(&Products{}).
		Select("products.id, products.name, products_suppliers.stock").
		Joins("left join products_suppliers on products_suppliers.product_id = products.id").
		Where("products_suppliers.supplier_id = ?", id).
		Distinct().
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(db *gorm.DB, id uint32) (Products, error) {
	var product Products
	err := db.First(&product, id).Error
	return product, err
}

func CreateProduct(db *gorm.DB, product *Products) error {
	return db.Create(product).Error
}

func UpdateProduct(db *gorm.DB, product *Products, id uint32) error {
	db.Updates(&product)
	return nil
}

func DeleteProduct(db *gorm.DB, id uint32) error {
	tx := db.Delete(&Products{}, id)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

func GetAttributeOfProducts(db *gorm.DB, attribute string) ([]string, error) {
	var result []string
	err := db.Model(&Products{}).Distinct(attribute).Pluck(attribute, &result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
