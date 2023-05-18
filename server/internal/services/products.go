package services

import (
	"server/db"
	"server/internal/models"
	"server/internal/utils"
	"strconv"
)

func GetSuppliersByProductID(id string) ([]map[string]interface{}, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return models.GetSuppliersByProductID(db.DB, uint32(idInt))
}

func GetProducts(query *models.ProductsQueryParams) ([]models.ProductsWithSuppliers, uint32, error) {
	products, total, err := models.GetProducts(db.DB, query)
	if err != nil {
		return nil, 0, err
	}

	var productsWithSuppliers []models.ProductsWithSuppliers
	for _, product := range products {
		suppliers, err := models.GetSuppliersByProductID(db.DB, product.ID)
		if err != nil {
			return nil, 0, err
		}
		productsWithSuppliers = append(productsWithSuppliers, models.ProductsWithSuppliers{
			Products:  product,
			Suppliers: suppliers,
		})
	}
	return productsWithSuppliers, total, nil
}

func GetProductByID(id uint32) (models.Products, error) {
	return models.GetProductByID(db.DB, id)
}

func CreateProduct(product *models.Products) error {
	product.Sku = utils.SkuGenerator(product)
	return models.CreateProduct(db.DB, product)
}

func UpdateProduct(product *models.Products, id uint32) error {
	_, err := models.GetProductByID(db.DB, id)
	if err != nil {
		return err
	}
	product.ID = id
	return models.UpdateProduct(db.DB, product, id)
}

func DeleteProduct(id uint32) error {
	return models.DeleteProduct(db.DB, id)
}

func GetAttributeOfProducts(attribute string) ([]map[string]interface{}, error) {
	values, err := models.GetAttributeOfProducts(db.DB, attribute)
	if err != nil {
		return make([]map[string]interface{}, 0), err
	}
	var attributeList []map[string]interface{}
	for _, value := range values {
		attributeList = append(attributeList, map[string]interface{}{
			"name": value,
		})
	}
	return attributeList, nil
}
