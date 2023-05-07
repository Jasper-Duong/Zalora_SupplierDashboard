package services

import (
	"net/url"
	"server/db"
	"server/internal/models"
)

func GetProducts(limit, offset int, params url.Values) ([]models.Products, int64, error) {
	products := []models.Products{}
	var total int64 = 0
	err := db.DB.Model(&models.Products{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.DB.Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
