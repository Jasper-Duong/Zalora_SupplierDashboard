package services

import (
	"server/db"
	"server/internal/models"
	"strconv"
)

func GetProducts(query *models.QueryParams) ([]models.Products, int64, error) {
	return models.SelectProducts(db.DB, query)
}

func GetProductByID(id string) (models.Products, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.Products{}, err
	}
	return models.SelectProductByID(db.DB, idInt)
}

func CreateProduct(product *models.Products) error {
	return models.CreateProduct(db.DB, product)
}

func UpdateProduct(product *models.Products, id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return models.UpdateProduct(db.DB, product, idInt)
}

func DeleteProduct(id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return models.DeleteProduct(db.DB, idInt)
}

func GetAttributeOfProducts(attribute string) ([]map[string]interface{}, error) {
	return models.GetAttributeOfProducts(db.DB, attribute)
}
