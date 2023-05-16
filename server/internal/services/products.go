package services

import (
	"server/db"
	"server/internal/models"
	"strconv"

	"server/utils"
)

func GetSuppliersByProductID(id string) ([]map[string]interface{}, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return models.GetSuppliersByProductID(db.DB, uint32(idInt))
}

func GetProducts(query *models.QueryParams) ([]models.ProductsWithSuppliers, uint32, error) {
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

func GetProductByID(id string) (models.Products, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.Products{}, err
	}
	return models.GetProductByID(db.DB, idInt)
}

func CreateProduct(product *models.Products) error {
	product.Sku = utils.SkuGenerator(product)
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
	/*err = models.DeleteStockByProductID(db.DB, uint32(idInt))
	if err != nil {
		return err
	}*/
	return models.DeleteProduct(db.DB, idInt)
}

func GetAttributeOfProducts(attribute string) ([]map[string]interface{}, error) {
	values, err := models.GetAttributeOfProducts(db.DB, attribute)
	if err != nil {
		return make([]map[string]interface{}, 0), err
	}
	var attributeList []map[string]interface{}
	for _, value := range values {
		attributeList = append(attributeList, map[string]interface{}{
			"text":  value,
			"value": value,
		})
	}
	return attributeList, nil
}
