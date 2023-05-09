package services

import (
	"fmt"
	"server/db"
	"server/internal/models"
	"strconv"
)

func GetProductStocks(id string) ([]map[string]interface{}, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	stocks, err := models.SelectProductStocks(db.DB, idInt)
	var stockMap []map[string]interface{}
	for _, stock := range stocks {
		stockMap = append(stockMap, map[string]interface{}{
			"id":    stock.SupplierID,
			"stock": stock.Stock,
		})
	}

	return stockMap, err
}

func CreateProductStock(stock map[string]interface{}, id string) error {
	product_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	supplier_id, ok := stock["id"].(float64)
	if !ok {
		return fmt.Errorf("id is not an integer")
	}

	stockValue, ok := stock["stock"].(float64)
	if !ok {
		return fmt.Errorf("stock is not an integer")
	}

	if err != nil {
		return err
	}

	stockObj := models.ProductsSuppliers{
		ProductID:  uint32(product_id),
		SupplierID: uint32(supplier_id),
		Stock:      uint32(stockValue),
	}

	return models.CreateProductStock(db.DB, &stockObj)
}

func UpdateProductStock(stock map[string]interface{}, product_id string, supplier_id string) error {
	product_id_int, err := strconv.Atoi(product_id)
	if err != nil {
		return err
	}

	supplier_id_int, err := strconv.Atoi(supplier_id)
	if err != nil {
		return err
	}

	stockInt, err := strconv.Atoi(stock["stock"].(string))
	if err != nil {
		return err
	}

	return models.UpdateProductStock(db.DB, uint32(product_id_int), uint32(supplier_id_int), uint32(stockInt))
}

func DeleteProductStock(product_id string, supplier_id string) error {
	product_id_int, err := strconv.Atoi(product_id)
	if err != nil {
		return err
	}

	supplier_id_int, err := strconv.Atoi(supplier_id)
	if err != nil {
		return err
	}

	return models.DeleteProductStock(db.DB, uint32(product_id_int), uint32(supplier_id_int))
}
