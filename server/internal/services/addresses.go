package services

import (
	"server/db"
	"server/internal/models"
)

func CreateAddress(data *models.AddressCreate) error {
	return models.CreateAddress(db.DB, data)
}

func UpdateAddress(data *models.AddressUpdate, id int) error {
	return models.UpdateAddress(db.DB, data, id)
}

func DeleteAddress(id int) error {
	return models.DeleteAddress(db.DB, id)
}
