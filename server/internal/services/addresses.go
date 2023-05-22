package services

import (
	"server/internal/db"
	"server/internal/models"
)

var CreateAddress = func(data *models.AddressCreate) error {
	if _, err := models.GetSupplierByID(db.DB, data.SupplierID); err != nil {
		return err
	}
	return models.CreateAddress(db.DB, data)
}

var UpdateAddress = func(data *models.AddressUpdate, id uint32) error {
	return models.UpdateAddress(db.DB, data, id)
}

var DeleteAddress = func(id uint32) error {
	if _, err := models.GetAddressByID(db.DB, id); err != nil {
		return err
	}
	return models.DeleteAddress(db.DB, id)
}
