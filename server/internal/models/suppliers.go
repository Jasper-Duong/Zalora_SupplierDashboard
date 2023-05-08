package models

import (
	"reflect"

	"gorm.io/gorm"
)

type Suppliers struct {
	ID            uint32 `gorm:"primary_key"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	ContactNumber string `json:"contact_number"`
	Status        bool   `gorm:"default:true"`
	Stock         uint32 `gorm:"default:0"`
}

type QueryParam struct {
	Page  int      `form:"page"`
	Limit int      `form:"limit"`
	Name  []string `form:"name[]"`
	Email []string `form:"email[]"`
}

func SelectSuppliers(db *gorm.DB, query *QueryParam) ([]Suppliers, error) {
	var suppliers []Suppliers

	db = db.Offset((query.Page - 1) * query.Limit).Limit(query.Limit)
	db = db.Where("name IN (?)", query.Name)

	err := db.Find(&suppliers).Error
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func field(value reflect.Value) {
	panic("unimplemented")
}

func CreateSupplier(db *gorm.DB, supplier *Suppliers) error {
	if err := db.Create(&supplier).Error; err != nil {
		return err
	}
	return nil
}
