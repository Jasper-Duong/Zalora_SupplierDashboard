package services

import (
	"server/db"
	"server/internal/models"
)

func CreateAddress(data *models.AddressCreate) error {
	return models.CreateAddress(db.DB, data)
}

func UpdateAddress(data *models.AddressUpdate, id uint32) error {
	return models.UpdateAddress(db.DB, data, id)
}

func DeleteAddress(id uint32) error {
	return models.DeleteAddress(db.DB, id)
}
