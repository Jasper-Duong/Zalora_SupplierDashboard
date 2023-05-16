package models

import (
	"reflect"

	"gorm.io/gorm"
)

type Products struct {
	ID     uint32 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name   string `gorm:"column:name;not null" json:"name" binding:"required"`
	Brand  string `gorm:"column:brand;not null" json:"brand" binding:"required"`
	Sku    string `gorm:"column:SKU;not null" json:"sku"`
	Size   string `gorm:"column:size;not null" json:"size" binding:"required"`
	Color  string `gorm:"column:color;not null" json:"color" binding:"required"`
	Status bool   `gorm:"column:status;not null" json:"status" binding:"required"`
	Stock  uint32 `gorm:"column:stock;default:0" json:"stock"`
}

type ProductsWithSuppliers struct {
	Products
	Suppliers []map[string]interface{} `json:"suppliers"`
}

type QueryParams struct {
	Page   int      `form:"page"`
	Limit  int      `form:"limit"`
	Brand  []string `form:"brand[]"`
	Size   []string `form:"size[]"`
	Color  []string `form:"color[]"`
	Status []string `form:"status[]"`
}

func GetProducts(db *gorm.DB, query *QueryParams) ([]Products, uint32, error) {
	var products []Products

	q := reflect.TypeOf(query).Elem()
	for i := 0; i < q.NumField(); i++ {
		field := q.Field(i)
		name := field.Name
		value := reflect.ValueOf(query).Elem().FieldByName(name)
		if name == "Page" || name == "Limit" {
			continue
		}
		if value.Len() > 0 {
			db = db.Where(name+" in ?", value.Interface())
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
	//total := uint32(len(products))
	//products = products[query.Page-1 : query.Page-1+query.Limit]

	return products, uint32(total), nil
}

func GetProductsBySupplierID(db *gorm.DB, id uint32) ([]map[string]interface{}, error) {
	var products []map[string]interface{} = make([]map[string]interface{}, 0)
	err := db.Model(&Products{}).
		Select("products.id, products.name").
		Joins("left join products_suppliers on products_suppliers.product_id = products.id").
		Where("products_suppliers.supplier_id = ?", id).
		Distinct().
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(db *gorm.DB, id int) (Products, error) {
	var product Products
	err := db.First(&product, id).Error
	if err != nil {
		return Products{}, err
	}
	return product, nil
}

func CreateProduct(db *gorm.DB, product *Products) error {
	return db.Create(product).Error
}

func UpdateProduct(db *gorm.DB, product *Products, id int) error {
	tx := db.Model(&product).Where("id = ?", id).Updates(product)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

func DeleteProduct(db *gorm.DB, id int) error {
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
