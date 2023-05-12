package models

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type Suppliers struct {
	ID            uint64 `gorm:"column:id;primaryKey;autoIncrement" `
	Name          string `gorm:"column:name;not null;unique" json:"name" validate:"required"`
	Email         string `gorm:"column:email;not null;unique" json:"email" validate:"required,email"`
	ContactNumber string `gorm:"column:contact_number;not null" json:"contact_number" validate:"required,e164"`
	Status        bool   `gorm:"column:status;default:true"`
	Stock         uint32 `gorm:"column:stock;default:0"`
}

type SuppliersInfo struct {
	ID   uint   `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

type QueryParam struct {
	Page           int      `form:"page"`
	Limit          int      `form:"limit"`
	Name           []string `form:"name[]"`
	Email          []string `form:"email[]"`
	Contact_number []string `form:"contact_number[]"`
}

func GetSuppliers(db *gorm.DB, query *QueryParam) ([]Suppliers, int64, error) {
	var suppliers []Suppliers

	db = db.Offset((query.Page - 1) * query.Limit).Limit(query.Limit)

	t := reflect.TypeOf(query).Elem()

	for i := 2; i < t.NumField(); i++ {
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

func CreateSupplier(db *gorm.DB, supplier *Suppliers) (int, error) {
	if err := db.Create(&supplier).Error; err != nil {
		return http.StatusConflict, err
	}
	return http.StatusOK, nil
}

func UpdateSupplier(db *gorm.DB, supplier *Suppliers) (int, error) {
	if err := db.First(&supplier).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}

	if err := db.Updates(&supplier).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func DeleteSupplier(db *gorm.DB, id uint64) (int, error) {
	if err := db.First(&Suppliers{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}
	if err := db.Delete(&Suppliers{}, id).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func GetSuppliersName(db *gorm.DB) ([]SuppliersInfo, error) {
	var suppliersInfo []SuppliersInfo
	if err := db.Model(&Suppliers{}).Select("id", "name").Find(&suppliersInfo).Error; err != nil {
		return []SuppliersInfo{}, err
	}
	return suppliersInfo, nil
}
