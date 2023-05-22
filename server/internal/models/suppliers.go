package models

import (
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type Suppliers struct {
	ID            uint32 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name          string `gorm:"column:name;not null;unique" json:"name" binding:"required"`
	Email         string `gorm:"column:email;not null;unique" json:"email" binding:"required,email"`
	ContactNumber string `gorm:"column:contact_number;not null" json:"contact_number" binding:"required"`
	Status        *bool  `gorm:"column:status;default:true" json:"status"`
	Stock         uint32 `gorm:"column:stock;default:0" json:"stock"`
}

type SupplierWithAddresses struct {
	Suppliers
	Addresses []map[string]interface{} `json:"addresses"`
}

type SuppliersQueryParam struct {
	Status         bool   `form:"status"`
	Page           int    `form:"page"`
	Limit          int    `form:"limit"`
	Name           string `form:"name[]"`
	Email          string `form:"email[]"`
	Contact_Number string `form:"contact_number[]"`
}

func GetSuppliers(db *gorm.DB, query *SuppliersQueryParam) ([]Suppliers, uint32, error) {
	var suppliers []Suppliers
	db = db.Where("status = ?", query.Status)

	t := reflect.TypeOf(query).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		//fieldType := field.Type
		name := field.Name
		value := reflect.ValueOf(query).Elem().Field(i)
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
	db.Model(&Suppliers{}).Count(&total)
	offset := (query.Page - 1) * query.Limit
	db = db.Offset(offset).Limit(query.Limit)
	err := db.Find(&suppliers).Error
	if err != nil {
		return nil, 0, err
	}

	return suppliers, uint32(total), nil
}

func GetSupplierByID(db *gorm.DB, ID uint32) (Suppliers, error) {
	var supplier Suppliers
	err := db.First(&supplier, ID).Error
	return supplier, err
}

func GetSuppliersByProductID(db *gorm.DB, id uint32) ([]map[string]interface{}, error) {
	var suppliers []map[string]interface{} = make([]map[string]interface{}, 0)
	err := db.Model(Suppliers{}).
		Select("suppliers.id, suppliers.name, products_suppliers.stock").
		Joins("left join products_suppliers on products_suppliers.supplier_id = suppliers.id").
		Where("products_suppliers.product_id = ?", id).
		Distinct().
		Find(&suppliers).Error
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func CreateSupplier(db *gorm.DB, supplier *Suppliers) error {
	return db.Create(supplier).Error
}

func UpdateSupplier(db *gorm.DB, supplier *Suppliers) error {
	return db.Updates(&supplier).Error
}

func DeleteSupplier(db *gorm.DB, id uint32) error {
	tx := db.Delete(&Suppliers{}, id)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func GetSuppliersAttribute(db *gorm.DB, attribute string) ([]map[string]interface{}, error) {
	var suppliers []map[string]interface{}
	if err := db.Model(&Suppliers{}).Select("id", attribute).Find(&suppliers).Error; err != nil {
		return make([]map[string]interface{}, 0), err
	}
	return suppliers, nil
}
