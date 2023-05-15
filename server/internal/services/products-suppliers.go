package services

import (
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

func GetSupplierStocks(id string) ([]map[string]interface{}, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	stocks, err := models.SelectSupplierStocks(db.DB, idInt)
	var stockMap []map[string]interface{}
	for _, stock := range stocks {
		stockMap = append(stockMap, map[string]interface{}{
			"id":    stock.ProductID,
			"stock": stock.Stock,
		})
	}

	return stockMap, err
}

func CreateStock(stock models.StockRequest) error {
	product_supplier := &models.ProductsSuppliers{
		ProductID:  stock.ProductID,
		SupplierID: stock.SupplierID,
		Stock:      stock.Stock,
	}
	return models.CreateStock(db.DB, product_supplier)
}

func UpdateStock(stock models.StockRequest) error {
	return models.UpdateStock(db.DB, stock.ProductID, stock.SupplierID, stock.Stock)
}

func DeleteStock(product_id string, supplier_id string) error {
	product_id_int, err := strconv.Atoi(product_id)
	if err != nil {
		return err
	}

	supplier_id_int, err := strconv.Atoi(supplier_id)
	if err != nil {
		return err
	}

	return models.DeleteStock(db.DB, uint32(product_id_int), uint32(supplier_id_int))
}
