package models

import (
	"errors"

	"gorm.io/gorm"
)

type Suppliers struct {
	ID            uint32 `gorm:"primary_key" `
	Name          string `gorm:"not null;unique" json:"name" validate:"required"`
	Email         string `gorm:"not null;unique" json:"email" validate:"required,email"`
	ContactNumber string `gorm:"not null" json:"contact_number" validate:"required,e164"`
	Status        bool   `gorm:"default:true"`
	Stock         uint32 `gorm:"default:0"`
}

type QueryParam struct {
	Page  int      `form:"page"`
	Limit int      `form:"limit"`
	Name  []string `form:"name[]"`
	Email []string `form:"email[]"`
}

func SelectSuppliers(db *gorm.DB, query *QueryParam) ([]Suppliers, int64, error) {
	var suppliers []Suppliers

	db = db.Offset((query.Page - 1) * query.Limit).Limit(query.Limit)
	db = db.Where("name IN (?)", query.Name)
	//db = db.Where("email IN (?)", query.Email)

	/*
		    value := reflect.ValueOf(QueryParam{})
		    for i := 0; i < value.NumField(); i++ {
		        field := value.Field(i)
				fieldName := value.Type().Field(i).Name
				fieldType := value.Type().Field(i)
				if fieldType.Type.Kind() == reflect.Slice {
					db = db.Where("(?) IN (?)", fieldName.ToLower(), )
				}
		    }
			"*/

	if err := db.Find(&suppliers).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := db.Model(&Suppliers{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return suppliers, total, nil
}

func CreateSupplier(db *gorm.DB, supplier *Suppliers) (int, error) {
	if err := db.Create(&supplier).Error; err != nil {
		return 409, err
	}
	return 200, nil
}

func UpdateSupplier(db *gorm.DB, supplier *Suppliers) (int, error) {
	if err := db.First(&supplier).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}

	if err := db.Updates(&supplier).Error; err != nil {
		return 500, err
	}
	return 200, nil
}
