package models

import (
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type Suppliers struct {
	ID            uint64 `gorm:"column:id;primaryKey;autoIncrement" `
	Name          string `gorm:"column:name;not null;unique" json:"name" binding:"required"`
	Email         string `gorm:"column:email;not null;unique" json:"email" binding:"required,email"`
	ContactNumber string `gorm:"column:contact_number;not null" json:"contact_number" binding:"required,e164"`
	Status        bool   `gorm:"column:status;default:true"`
	Stock         uint32 `gorm:"column:stock;default:0"`
}

type SuppliersInfo struct {
	ID   uint   `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

type SuppliersQueryParam struct {
	Page           int      `form:"page"`
	Limit          int      `form:"limit"`
	Name           []string `form:"name[]"`
	Email          []string `form:"email[]"`
	Contact_number []string `form:"contact_number[]"`
}

func GetSuppliers(db *gorm.DB, query *SuppliersQueryParam) ([]Suppliers, int64, error) {
	var suppliers []Suppliers

	db = db.Offset((query.Page - 1) * query.Limit).Limit(query.Limit)

	t := reflect.TypeOf(query).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		//fieldType := field.Type
		fieldName := field.Name
		fieldValue := reflect.ValueOf(query).Elem().Field(i)

		if fieldValue.Kind() == reflect.Slice {
			length := fieldValue.Len()
			if length > 0 {
				db = db.Where(strings.ToLower(fieldName)+" IN ?", fieldValue.Interface())
			}
		}
	}

	if err := db.Find(&suppliers).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := db.Model(&Suppliers{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return suppliers, total, nil
}

func CreateSupplier(db *gorm.DB, supplier *Suppliers) error {
	if err := db.Create(&supplier).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSupplier(db *gorm.DB, supplier *Suppliers) error {
	tx := db.Updates(&supplier)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DeleteSupplier(db *gorm.DB, id uint64) error {
	tx := db.Delete(&Suppliers{}, id)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func GetSuppliersName(db *gorm.DB) ([]SuppliersInfo, error) {
	var suppliersInfo []SuppliersInfo
	if err := db.Model(&Suppliers{}).Select("id", "name").Find(&suppliersInfo).Error; err != nil {
		return []SuppliersInfo{}, err
	}
	return suppliersInfo, nil
}
