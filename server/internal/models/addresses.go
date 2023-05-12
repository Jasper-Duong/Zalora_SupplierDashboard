package models

import (
	"gorm.io/gorm"
)

type Addresses struct {
	ID          uint32    `gorm:"column:id;primaryKey;autoIncrement"`
	SupplierID  uint32    `gorm:"column:supplier_id;not null"`
	Type        string    `gorm:"column:type; default:office"`
	AddressInfo string    `gorm:"column:address_info; not null"`
	Suppliers   Suppliers `gorm:"foreignKey:SupplierID; constraint:OnDelete: CASCADE"`
}

func (Addresses) TableName() string { return "addresses" }

type AddressCreate struct {
	SupplierID  uint32 `json:"supplier_id" gorm:"column:supplier_id;" binding:"required"`
	Type        string `json:"type" gorm:"column:type;" binding:"required"`
	AddressInfo string `json:"address_info" gorm:"column:address_info;" binding:"required"`
}

type AddressUpdate struct {
	Type        string `json:"type" gorm:"column:type;" binding:"required"`
	AddressInfo string `json:"address_info" gorm:"column:address_info;" binding:"required"`
}

func (AddressUpdate) TableName() string { return "addresses" }

func (AddressCreate) TableName() string { return "addresses" }

func GetAddresses() Addresses {
	address := Addresses{
		ID:          1,
		SupplierID:  1,
		Type:        "office",
		AddressInfo: "123 Street",
	}
	return address
}

func CreateAddress(db *gorm.DB, data *AddressCreate) error {
	return db.Create(&data).Error
}

func UpdateAddress(db *gorm.DB, data *AddressUpdate, id int) error {
	tx := db.Where("id = ?", id).Updates(&data)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

func DeleteAddress(db *gorm.DB, id int) error {
	tx := db.Delete(&Addresses{}, id)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}
