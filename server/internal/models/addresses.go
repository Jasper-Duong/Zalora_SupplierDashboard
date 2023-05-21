package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Addresses struct {
	ID          uint32    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SupplierID  uint32    `gorm:"column:supplier_id;not null" json:"supplier_id"`
	Type        string    `gorm:"column:type; default:office" json:"type"`
	AddressInfo string    `gorm:"column:address_info; not null" json:"address_info"`
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

type AddressesRepository struct {
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

func GetAddressByID(db *gorm.DB, id uint32) (Addresses, error) {
	var address Addresses
	err := db.First(&address, id).Error
	return address, err
}

func CreateAddress(db *gorm.DB, data *AddressCreate) error {
	return db.Create(&data).Error
}

func UpdateAddress(db *gorm.DB, data *AddressUpdate, id uint32) error {
	var old *AddressUpdate
	if err := db.First(&old, "id = ?", id).Error; err != nil {
		return err
	}
	fmt.Println(data)
	return db.Where("id = ?", id).Updates(&data).Error
}

func DeleteAddress(db *gorm.DB, id uint32) error {
	tx := db.Delete(&Addresses{}, id)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

func GetAddressesBySupplierID(db *gorm.DB, id uint32) ([]map[string]interface{}, error) {
	var addresses = make([]map[string]interface{}, 0)
	err := db.Model(Addresses{}).Select("id", "type", "address_info").Where("supplier_id = ?", id).Find(&addresses).Error
	return addresses, err
}
